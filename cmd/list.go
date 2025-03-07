package cmd

import (
	"fmt"
	"sort"

	"github.com/masamerc/sevp/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available selectors",
	Run: func(cmd *cobra.Command, args []string) {
		selectorMap, err := internal.GetSelectors()
		internal.FailOnError("Error getting selectors", err)

		var selectorSlice []string

		for k := range selectorMap {
			selectorSlice = append(selectorSlice, k)
		}

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
