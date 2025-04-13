package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/app"
	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(viewCmd)
}

// viewCmd displays the details of a selector.
var viewCmd = &cobra.Command{
	Use:   "view <selector>",
	Short: "Check the possible values of a selector",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run:   runView,
}

// runView executes the view command, displaying the details of a selector.
func runView(cmd *cobra.Command, args []string) {
	var selector internal.Selector

	// Config found
	selectorMap, err := internal.ParseSelectorsFromConfig()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "Error getting selectors: %v\n", err)
		return
	}

	selectorChoice := args[0]

	if selectorChosen, ok := selectorMap[selectorChoice]; ok {
		// If the selector is a external config selector,
		// sevp will attempt to read and parse the external config file.
		if selectorChosen.ReadConfig {
			selector, err = selectorChosen.IntoExternalConfigSelector()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "Failed to parse selectors: %v\n", err)
				return
			}
		} else {
			selector = selectorChosen
		}
	} else {
		fmt.Fprintf(cmd.OutOrStderr(), "Selector not found: %v\n", err)
		return
	}

	// Read the content of the selector
	targetVar, possibleValues, err := selector.Read()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "Failed to parse selectors: %v\n", err)
		return
	}

	// Some styling for the stdout
	purpleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightPurple))
	greenStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightGreen))

	// Display
	fmt.Fprintf(cmd.OutOrStdout(), "\ntarget environment variable:\n  %s\n", purpleStyle.Render(targetVar))
	fmt.Fprintf(cmd.OutOrStdout(), "\npossible values:\n")

	for _, v := range possibleValues {
		fmt.Fprintf(cmd.OutOrStdout(), "  - %s\n", greenStyle.Render(v))
	}
}
