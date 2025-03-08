package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Selector interface {
	Read() (string, []string, error)
}

type configSelectorMap map[string]*ConfigSelector

type ConfigSelector struct {
	Name           string
	ReadConfig     bool
	TargetVar      string
	PossibleValues []string
}

func (s *ConfigSelector) Read() (string, []string, error) {
	return s.TargetVar, s.PossibleValues, nil
}

// IntoExternalProviderSelector converts the config selector into a external provider selector
// For example, the external provider selector for AWS will read profiles set in the AWS config (~/.aws/config)
//
// Returns:
//   - Selector: The external provider selector
func (s *ConfigSelector) IntoExternalProviderSelector() (Selector, error) {
	switch s.Name {
	case "aws":
		return &AWSProfileSelector{}, nil
	default:
		return nil, fmt.Errorf("the external config provider is not supported for selector %s", s.Name)
	}
}

// FromConfig creates a config selector from the viper config
//
// Args:
//   - name: The name of the selector
//
// Returns:
//   - *ConfigSelector: The config selector
func FromConfig(name string) (*ConfigSelector, error) {
	readConfig := viper.GetBool(name + ".read_config")
	targetVar := viper.GetString(name + ".target_var")
	possibleValues := viper.GetStringSlice(name + ".possible_values")

	if (targetVar == "" || len(possibleValues) == 0) && !readConfig {
		return nil, fmt.Errorf(
			"invalid selector: %s - either the selector is not in the config or the `target_var` or `possible_values` is not set for the selector",
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
// Viper looks for the config in $HOME/.config/sevp.toml and $HOME/sevp.toml.
// $Home/.config/sevp.toml takes precedence over $HOME/sevp.toml if both exist.
//
// Returns:
//   - error: An error if the config file cannot be read
func ParseConfig() error {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
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
	slog.Debug("Config file read successfully", "path", viper.ConfigFileUsed())
	return nil
}

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
