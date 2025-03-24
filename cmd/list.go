package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/internal"
)

// list command will list all selectors in the config.
// list command will only work if a config is found
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available selectors",
	Run: func(cmd *cobra.Command, args []string) {
		selectorMap, err := internal.GetSelectors()
		failOnError("Error getting selectors", err)

		var selectorSlice []string

		for k := range selectorMap {
			selectorSlice = append(selectorSlice, k)
		}

		// sort the selectors since the map order is not guaranteed
		sorted := sort.StringSlice(selectorSlice)
		sort.Sort(sorted)

		for _, s := range sorted {
			fmt.Println(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
