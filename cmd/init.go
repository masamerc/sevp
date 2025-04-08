package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/internal"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmd prints out a shell hook for the supported shells.
var initCmd = &cobra.Command{
	Use:       "init <shell>",
	Short:     fmt.Sprintf("Prints out a shell hook for the specified shell. Supported shells: %v", internal.SupportedShells),
	ValidArgs: internal.SupportedShells,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:       runInit,
}

// shellToHook maps the shell name to the corresponding hook function.
var shellToHook = map[string]string{
	"bash": internal.BashHook,
	"zsh":  internal.ZshHook,
}

// runInit executes the init command, printing the shell hook for the specified shell.
func runInit(cmd *cobra.Command, args []string) {
	fmt.Fprintln(cmd.OutOrStdout(), shellToHook[args[0]])
}
