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
	Run: runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	var selector internal.Selector

	if viper.ConfigFileUsed() != "" {
		// config route
		if len(args) == 1 {
			selectorChoice := args[0]

			// config parse
			selecotrSection, err := internal.FromConfig(selectorChoice)
			failOnError("Failed to parse selectors", err)
			selector = selecotrSection

		} else {
			defaultSelector := viper.GetString("default")
			slog.Debug("default selector: " + defaultSelector)

			selecotrSection, err := internal.FromConfig(defaultSelector)
			failOnError("Failed to parse selectors", err)

			if selecotrSection.ReadConfig && defaultSelector == "aws" {
				selector = internal.NewAWSProfileSelector()
			} else {
				if selecotrSection.TargetVar == "" || len(selecotrSection.PossibleValues) == 0 {
					failOnError("Error getting selectors", fmt.Errorf("missing target_var or possible_values"))
				}
				selector = selecotrSection
			}
		}
	} else {
		// no config -> aws config mode
		selector = internal.NewAWSProfileSelector()
	}

	startApp(selector)
}

func startApp(s internal.Selector) {
	targetVar, possibleValues, err := s.Read()
	failOnError("Failed to parse selectors", err)
	app := app.NewApp(possibleValues, targetVar)
	err = app.Run()
	failOnError("Failed to run app", err)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
			failOnError("AWS config file not found", err)
		} else {
			failOnError("Error checking AWS config file", err)
		}
	}
}

// Fail logs the message and exits the program with a non-zero status code.
//
// Parameters:
//   - msg: The error message to log.
func fail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// failOnError logs the error message and exits the program if an error is encountered.
//
// Parameters:
//   - msg: A descriptive message to log when an error occurs.
//   - err: The error to log and handle.
func failOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err)
		os.Exit(1)
	}
}
