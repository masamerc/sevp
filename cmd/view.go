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

// view command will display the details of a selector.
// if no config is found, it will default to AWS profiles.
var viewCmd = &cobra.Command{
	Use:   "view <selector>",
	Short: "Check the possible values of a selector",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var selector internal.Selector

		if viper.ConfigFileUsed() == "" {
			// no config -> default to AWS
			selector = internal.NewAWSProfileSelector()
		} else {
			// config found
			selectorMap, err := internal.GetSelectors()
			failOnError("Error getting selectors", err)

			selectorChoice := args[0]

			if selectorChosen, ok := selectorMap[selectorChoice]; ok {
				if selectorChosen.ReadConfig {
					selector, err = selectorChosen.IntoExternalProviderSelector()
					failOnError("Failed to parse selectors", err)
				} else {
					selector = selectorChosen
				}
			} else {
				failOnError("Selector not found", fmt.Errorf("selector %s not found", selectorChoice))
			}
		}

		// read the content of the selector
		targetVar, possibleValues, err := selector.Read()
		failOnError("Failed to parse selectors", err)

		// some styling for the stdout
		purpleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightPurple))
		greenSytle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightGreen))

		// display
		fmt.Printf("\ntarget environment variable:\n  %s\n", purpleStyle.Render(targetVar))
		fmt.Println("\npossible values:")
		for _, v := range possibleValues {
			fmt.Printf("  - %s\n", greenSytle.Render(v))
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
