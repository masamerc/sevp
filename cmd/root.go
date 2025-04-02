package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/app"
	"github.com/masamerc/sevp/internal"
)

var rootCmd = &cobra.Command{
	Use:     "sevp [command]",
	Version: "0.1.0",
	Short:   "sevp: pick and switch environement variables.",
	Args: cobra.MatchAll(
		cobra.MinimumNArgs(0),
		cobra.MaximumNArgs(1),
	),
	RunE: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	selector, err := internal.InitSelector(args)

	if err != nil {
		return err
	}

	targetVar, possibleValues, err := selector.Read()

	if err != nil {
		return fmt.Errorf("failed to parse selectors: %w", err)
	}

	app := app.NewApp(possibleValues, targetVar)

	if err := app.Run(); err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if err := internal.InitConfig(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	})
}
