/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/masamerc/sevp/app"
	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(viewCmd)
}

// view command will display the details of a selector.
// if no config is found, it will default to AWS profiles.
var viewCmd = &cobra.Command{
	Use:   "view <selector>",
	Short: "Check the possible values of a selector",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run:   runView,
}

func runView(cmd *cobra.Command, args []string) {
	var selector internal.Selector

	if viper.ConfigFileUsed() == "" {
		// no config -> default to AWS
		selector = internal.NewAWSProfileSelector()
	} else {
		// config found
		selectorMap, err := internal.GetSelectors()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error getting selectors: %v\n", err)
			return
		}

		selectorChoice := args[0]

		if selectorChosen, ok := selectorMap[selectorChoice]; ok {
			if selectorChosen.ReadConfig {
				selector, err = selectorChosen.IntoExternalProviderSelector()
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), "Failed to parse selectors: %v\n", err)
					return
				}
			} else {
				selector = selectorChosen
			}
		} else {
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Selector not found: %v\n", err)
				return
			}
		}
	}

	// read the content of the selector
	targetVar, possibleValues, err := selector.Read()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "Failed to parse selectors: %v\n", err)
		return
	}

	// some styling for the stdout
	purpleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightPurple))
	greenSytle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightGreen))

	// display
	fmt.Fprintf(cmd.OutOrStdout(), "\ntarget environment variable:\n  %s\n", purpleStyle.Render(targetVar))
	fmt.Fprintf(cmd.OutOrStdout(), "\npossible values:\n")

	for _, v := range possibleValues {
		fmt.Fprintf(cmd.OutOrStdout(), "  - %s\n", greenSytle.Render(v))
	}
}
