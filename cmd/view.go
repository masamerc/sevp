/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/masamerc/sevp/internal"
	"github.com/masamerc/sevp/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viewCmd = &cobra.Command{
	Use:   "view <selector>",
	Short: "Check the possible values of a selector",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var selector internal.Selector

		if viper.ConfigFileUsed() == "" {
			// no config
			selector = internal.NewAWSProfileSelector()
		} else {
			// config found
			selectorMap, err := internal.GetSelectors()
			internal.FailOnError("Error getting selectors", err)

			selectorChoice := args[0]

			if selectorChosen, ok := selectorMap[selectorChoice]; ok {
				if selectorChosen.ReadConfig {
					selector, err = selectorChosen.IntoExternalProviderSelector()
					internal.FailOnError("Failed to parse selectors", err)
				} else {
					selector = selectorChosen
				}
			} else {
				internal.FailOnError("Selector not found", fmt.Errorf("selector %s not found", selectorChoice))
			}
		}
		displaySelectorInfo(selector)
	},
}

func displaySelectorInfo(s internal.Selector) {
	targetVar, possibleValues, err := s.Read()
	internal.FailOnError("Failed to parse selectors", err)

	purpleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightPurple))
	greenSytle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightGreen))

	fmt.Printf("\ntarget environment variable:\n  %s\n", purpleStyle.Render(targetVar))
	fmt.Println("\npossible values:")

	for _, v := range possibleValues {
		fmt.Printf("  - %s\n", greenSytle.Render(v))
	}
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
