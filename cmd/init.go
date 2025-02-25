package cmd

import (
	"fmt"
	"os"

	"github.com/masamerc/sevp/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:       "init <shell>",
	Short:     fmt.Sprintf("Prints out a shell-init for the input shell. Supported shells: %v", internal.SupportedShells),
	ValidArgs: internal.SupportedShells,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			fmt.Println(internal.Bash{}.Hook())
		case "zsh":
			fmt.Println(internal.Zsh{}.Hook())
		case "fish":
			fmt.Println(internal.Fish{}.Hook())
		case "nu":
			fmt.Println(internal.Nu{}.Hook())
		default:
			fmt.Fprintf(os.Stderr, "Error: enter a valid shell: %v\n", internal.SupportedShells)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
