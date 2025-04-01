package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// init command will print out a shell hook for the supported shells.
var initCmd = &cobra.Command{
	Use:       "init <shell>",
	Short:     fmt.Sprintf("Prints out a shell-init for the input shell. Supported shells: %v", internal.SupportedShells),
	ValidArgs: internal.SupportedShells,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:       runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Error: enter a valid shell: %v\n", internal.SupportedShells)
		os.Exit(1)
	}
	switch args[0] {
	case "bash":
		fmt.Fprintln(os.Stdout, internal.BashHook)
	case "zsh":
		fmt.Fprintln(os.Stdout, internal.ZshHook)
	default:
		fmt.Fprintf(os.Stderr, "Error: enter a valid shell: %v\n", internal.SupportedShells)
		os.Exit(1)
	}
}
