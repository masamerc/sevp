package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/masamerc/sevp/app"
	"github.com/masamerc/sevp/internal"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:     "sevp [command]",
	Version: "1.0.3",
	Short:   "sevp: simple environment variable picker",
	Args: cobra.MatchAll(
		cobra.MinimumNArgs(0),
		cobra.MaximumNArgs(1),
	),
	RunE: runRoot,
}

// runRoot acts as the main entry point for the entire CLI application.
func runRoot(cmd *cobra.Command, args []string) error {
	selector, err := internal.GetSelector(args)
	if err != nil {
		return err
	}

	targetVar, possibleValues, err := selector.Read()
	if err != nil {
		return err
	}

	app := app.NewApp(possibleValues, targetVar)

	if err := app.Run(); err != nil {
		return err
	}

	return nil
}

// Execute is the main entry point for the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// init initializes the CLI application by setting up the configuration.
func init() {
	cobra.OnInitialize(func() {
		internal.InitLogger()

		if err := internal.InitConfig(); err != nil {
			// If config is not found, create one with the default config content
			// and exit so the CLI can pick up the new config
			if err.Error() == "created default config" {
				fmt.Fprintln(os.Stderr, "Created default config: $HOME/.config/sevp.toml.")
				fmt.Fprintln(os.Stderr, "Try running sevp again or edit the config to your needs.")
				os.Exit(0)
			}

			// For other errors just exit
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	})
}
