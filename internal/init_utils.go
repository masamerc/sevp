package internal

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

// InitConfig initializes the configuration by setting up logging and reading the config file.
func InitConfig() error {
	// Initialize logger
	InitLogger()

	// Read in config
	if err := ParseConfig(); err != nil {
		slog.Debug("Error parsing config", "err", err)
		viper.SetDefault("default", "aws")
	}

	// Check for AWS config
	path, err := GetAWSConfigFile()
	if err != nil {
		slog.Debug("Error getting AWS config path", "err", err)
	}

	// Check if the AWS config file exists
	_, err = os.Stat(path)

	// If both AWS config and SEVP config are missing, return an error
	if err != nil && viper.ConfigFileUsed() == "" {
		if os.IsNotExist(err) {
			return fmt.Errorf("AWS config file not found: %w", err)
		}
		return fmt.Errorf("error checking AWS config file: %w", err)
	}
	return nil
}

// InitializeSelector sets up the appropriate selector based on the provided arguments and configuration.
func InitSelector(args []string) (Selector, error) {
	var selector Selector

	if viper.ConfigFileUsed() != "" {
		// Config route
		if len(args) == 1 {
			selectorChoice := args[0]

			// Config parse
			selectorSection, err := FromConfig(selectorChoice)
			if err != nil {
				return nil, fmt.Errorf("failed to parse selectors: %w", err)
			}
			selector = selectorSection

		} else {
			defaultSelector := viper.GetString("default")
			slog.Debug("default selector: " + defaultSelector)

			selectorSection, err := FromConfig(defaultSelector)
			if err != nil {
				return nil, fmt.Errorf("failed to parse selectors: %w", err)
			}

			if selectorSection.ReadConfig && defaultSelector == "aws" {
				selector = NewAWSProfileSelector()
			} else {
				if selectorSection.TargetVar == "" || len(selectorSection.PossibleValues) == 0 {
					return nil, fmt.Errorf("missing target_var or possible_values")
				}
				selector = selectorSection
			}
		}
	} else {
		// No config -> AWS config mode
		selector = NewAWSProfileSelector()
	}

	return selector, nil
}
