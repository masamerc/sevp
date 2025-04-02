package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

// list command will list all selectors in the config.
// only works if the config file is present.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available selectors",
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	selectorMap, err := internal.GetSelectors()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "Error getting selectors: %v\n", err)
		return
	}

	var selectorSlice []string

	for k := range selectorMap {
		selectorSlice = append(selectorSlice, k)
	}

	// sort the selectors since the map order is not guaranteed
	sorted := sort.StringSlice(selectorSlice)
	sort.Sort(sorted)

	for _, s := range sorted {
		fmt.Fprintln(cmd.OutOrStdout(), s)
	}
}
