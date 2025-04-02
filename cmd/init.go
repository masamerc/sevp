package cmd

import (
	"fmt"

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

var shellToHook = map[string]string{
	"bash": internal.BashHook,
	"zsh":  internal.ZshHook,
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Fprintf(cmd.OutOrStdout(), shellToHook[args[0]])

}
