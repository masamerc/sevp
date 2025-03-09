package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/masamerc/sevp/src/app"
	"github.com/masamerc/sevp/src/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "sevp [command]",
	Version: "0.1.0",
	Short:   "sevp: pick and switch environement variables.",
	Args: cobra.MatchAll(
		cobra.MinimumNArgs(0),
		cobra.MaximumNArgs(1),
	),
	Run: func(cmd *cobra.Command, args []string) {
		var selector internal.Selector

		if viper.ConfigFileUsed() != "" {
			// config route
			if len(args) == 1 {
				selectorChoice := args[0]

				// config parse
				selecotrSection, err := internal.FromConfig(selectorChoice)
				internal.FailOnError("Failed to parse selectors", err)
				selector = selecotrSection

			} else {
				defaultSelector := viper.GetString("default")
				slog.Debug("default selector: " + defaultSelector)

				selecotrSection, err := internal.FromConfig(defaultSelector)
				internal.FailOnError("Failed to parse selectors", err)

				if selecotrSection.ReadConfig && defaultSelector == "aws" {
					selector = internal.NewAWSProfileSelector()
				} else {
					if selecotrSection.TargetVar == "" || len(selecotrSection.PossibleValues) == 0 {
						internal.FailOnError("Error getting selectors", fmt.Errorf("missing target_var or possible_values"))
					}
					selector = selecotrSection
				}
			}
		} else {
			// no config -> aws config mode
			selector = internal.NewAWSProfileSelector()
		}

		startApp(selector)
	},
}

func startApp(s internal.Selector) {
	targetVar, possibleValues, err := s.Read()
	internal.FailOnError("Failed to parse selectors", err)
	app := app.NewApp(possibleValues, targetVar)
	err = app.Run()
	internal.FailOnError("Failed to run app", err)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// log settings
	internal.InitLogger()

	// read in config
	err := internal.ParseConfig()
	if err != nil {
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

	// if both aws config and sevp coinfig are missing, exit.
	// this is because if we have the aws config the aws selector will at least work,
	// and if we have the sevp config, the app will work without aws config
	if err != nil && viper.ConfigFileUsed() == "" {
		if os.IsNotExist(err) {
			internal.FailOnError("AWS config file not found", err)
		} else {
			internal.FailOnError("Error checking AWS config file", err)
		}
	}
}
