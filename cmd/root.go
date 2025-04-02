package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	selector, err := initializeSelector(args)
	if err != nil {
		return err
	}
	return startApp(selector)
}

func startApp(s internal.Selector) error {
	targetVar, possibleValues, err := s.Read()
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
		if err := initConfig(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	})
}

func initConfig() error {
	// log settings
	internal.InitLogger()

	// read in config
	if err := internal.ParseConfig(); err != nil {
		slog.Debug("Error parsing config", "err", err)
		viper.SetDefault("default", "aws")
	}

	// check for AWS config
	path, err := internal.GetAWSConfigFile()
	if err != nil {
		slog.Debug("Error getting AWS config path", "err", err)
	}

	// check if the AWS config file exists
	_, err = os.Stat(path)

	// if both aws config and sevp config are missing, return an error
	if err != nil && viper.ConfigFileUsed() == "" {
		if os.IsNotExist(err) {
			return fmt.Errorf("AWS config file not found: %w", err)
		}
		return fmt.Errorf("error checking AWS config file: %w", err)
	}
	return nil
}

func initializeSelector(args []string) (internal.Selector, error) {
	var selector internal.Selector

	if viper.ConfigFileUsed() != "" {
		// config route
		if len(args) == 1 {
			selectorChoice := args[0]

			// config parse
			selectorSection, err := internal.FromConfig(selectorChoice)
			if err != nil {
				return nil, fmt.Errorf("failed to parse selectors: %w", err)
			}
			selector = selectorSection

		} else {
			defaultSelector := viper.GetString("default")
			slog.Debug("default selector: " + defaultSelector)

			selectorSection, err := internal.FromConfig(defaultSelector)
			if err != nil {
				return nil, fmt.Errorf("failed to parse selectors: %w", err)
			}

			if selectorSection.ReadConfig && defaultSelector == "aws" {
				selector = internal.NewAWSProfileSelector()
			} else {
				if selectorSection.TargetVar == "" || len(selectorSection.PossibleValues) == 0 {
					return nil, fmt.Errorf("missing target_var or possible_values")
				}
				selector = selectorSection
			}
		}
	} else {
		// no config -> aws config mode
		selector = internal.NewAWSProfileSelector()
	}

	return selector, nil
}
