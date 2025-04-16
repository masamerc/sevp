package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/app"
	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(listCmd)

	// quite flag: quiet mode
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "hide target variables and their current values")
}

// listCmd lists all selectors in the config.
// It only works if the config file is present.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available selectors",
	Run:   runList,
}

// runList executes the list command, printing all available selectors.
func runList(cmd *cobra.Command, args []string) {
	selectorMap, err := internal.ParseSelectorsFromConfig()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "Error getting selectors: %v\n", err)
		return
	}

	var selectorSlice []string

	for s := range selectorMap {
		selectorSlice = append(selectorSlice, s)
	}

	// Sort the selectors since the map order is not guaranteed
	sorted := sort.StringSlice(selectorSlice)
	sort.Sort(sorted)

	// Quite mode will just print the selectors
	if quiet, _ := cmd.Flags().GetBool("quiet"); quiet {
		for _, s := range sorted {
			fmt.Fprintln(cmd.OutOrStdout(), s)
		}
		return
	}

	// Styling
	purpleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightPurple))
	greenStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(app.HexBrightGreen))

	// Calculate max width for padding;w
	maxWidth := 0
	for _, s := range sorted {
		if len(s) > maxWidth {
			maxWidth = len(s)
		}
	}
	maxWidth += 2

	for _, s := range sorted {
		currentSelector := selectorMap[s]
		currentTargetVar := currentSelector.TargetVar
		currentValue := os.Getenv(currentTargetVar)

		paddedName := fmt.Sprintf("%-*s", maxWidth, s) // left-aligned to width
		// nameStyled := purpleStyle.Render(paddedName)
		currentStyled := fmt.Sprintf(
			"(current: %v = %v)",
			purpleStyle.Render(currentTargetVar),
			greenStyle.Render(currentValue),
		)

		fmt.Fprintf(cmd.OutOrStdout(), "%s %s\n", paddedName, currentStyled)
	}
}
