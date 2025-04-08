package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// Selector is an interface that defines a method for reading configuration values.
type Selector interface {
	Read() (string, []string, error)
}

// configSelectorMap is a type alias for a map of string to ConfigSelector pointers.
type configSelectorMap map[string]*ConfigSelector

// ConfigSelector is a struct that defines a set of custom configuration options for a selector.
type ConfigSelector struct {
	Name           string
	ReadConfig     bool
	TargetVar      string
	PossibleValues []string
}

// Read is a method that reads the configuration values from the selector.
func (s *ConfigSelector) Read() (string, []string, error) {
	return s.TargetVar, s.PossibleValues, nil
}

// IntoExternalProviderSelector converts the config selector into a external provider selector
//
// For example, the external provider selector for AWS will read profiles set in the AWS config (~/.aws/config)
func (s *ConfigSelector) IntoExternalProviderSelector() (Selector, error) {
	switch s.Name {
	case "aws":
		return &AWSProfileSelector{}, nil
	default:
		return nil, fmt.Errorf("the external config provider is not supported for selector %s", s.Name)
	}
}

// FromConfig creates a config selector from the viper config
func FromConfig(name string) (*ConfigSelector, error) {
	readConfig := viper.GetBool(name + ".read_config")
	targetVar := viper.GetString(name + ".target_var")
	possibleValues := viper.GetStringSlice(name + ".possible_values")

	if (targetVar == "" || len(possibleValues) == 0) && !readConfig {
		return nil, fmt.Errorf(
			"invalid selector: %s - either the selector is not in the config, the `target_var` or `possible_values` is not set for the selector or the config file is not found",
			name,
		)
	}

	return &ConfigSelector{
		Name:           name,
		ReadConfig:     readConfig,
		TargetVar:      targetVar,
		PossibleValues: possibleValues,
	}, nil
}

// ParseConfig reads in the config file.
//
// Viper looks for the config in $HOME/.config/sevp.toml and $HOME/sevp.toml.
// $Home/.config/sevp.toml takes precedence over $HOME/sevp.toml if both exist.
func ParseConfig() error {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Set the config name and type
	viper.SetConfigName("sevp") // No dot here, Viper automatically adds extensions
	viper.SetConfigType("toml")

	// Add both possible config paths
	// the .config takes precedence over the home directory
	viper.AddConfigPath(path.Join(home, ".config")) // $HOME/.config/sevp.toml
	viper.AddConfigPath(home)                       // $HOME/sevp.toml

	// Read in the config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			slog.Debug("Config file not found")
		} else {
			slog.Debug("Error reading config", "err", err)
		}

		return err
	}

	if len(viper.AllSettings()) == 0 {
		return fmt.Errorf("config file is empty")
	}

	slog.Debug("Config file read successfully", "path", viper.ConfigFileUsed())
	return nil
}

// GetSelectors reads the config file and returns a map of selectors.
func GetSelectors() (configSelectorMap, error) {
	if viper.ConfigFileUsed() == "" {
		return nil, fmt.Errorf("no config file found")
	}

	selectors := make(configSelectorMap)
	topLevelKeysSet := make(map[string]struct{})

	for key := range viper.AllSettings() {
		if key == "default" {
			continue
		}
		// Extract the top-level key efficiently
		topLevelKey := strings.SplitN(key, ".", 2)[0]
		topLevelKeysSet[topLevelKey] = struct{}{} // Store unique top-level keys

		// Process selector configuration
		s, err := FromConfig(key)
		if err != nil {
			return nil, fmt.Errorf("error processing selector %s: %v", key, err)
		}
		selectors[key] = s
	}

	return selectors, nil
}
