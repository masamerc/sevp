package internal

import (
	_ "embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

//go:embed default_config.toml
var defaultConfig string

// Selector is an interface that defines a method for reading configuration values.
type Selector interface {
	Read() (string, []string, error)
}

// ConfigSelector is a struct that defines a set of custom configuration options for a selector.
type ConfigSelector struct {
	Name               string
	ReadExternalConfig bool
	TargetVar          string
	PossibleValues     []string
}

// Read is a method that reads the configuration values from the selector.
func (s *ConfigSelector) Read() (string, []string, error) {
	return s.TargetVar, s.PossibleValues, nil
}

// IntoExternalConfigSelector converts the config selector into a external provider selector
//
// For example, the external provider selector for AWS will read profiles set in the AWS config (~/.aws/config)
func (s *ConfigSelector) IntoExternalConfigSelector() (Selector, error) {
	return GetExternalConfigSelector(s.Name)
}

// FromConfig creates a config selector from the viper config
func FromConfig(name string) (*ConfigSelector, error) {
	readConfig := viper.GetBool(name + ".external_config")
	targetVar := viper.GetString(name + ".target_var")
	possibleValues := viper.GetStringSlice(name + ".possible_values")

	if (targetVar == "" || len(possibleValues) == 0) && !readConfig {
		return nil, fmt.Errorf(
			"invalid selector: %s - either the selector is not in the config, the `target_var` or `possible_values` is not set for the selector or the config file is not found",
			name,
		)
	}

	return &ConfigSelector{
		Name:               name,
		ReadExternalConfig: readConfig,
		TargetVar:          targetVar,
		PossibleValues:     possibleValues,
	}, nil
}

// InitConfig initializes the configuration by  reading the config file.
func InitConfig() error {
	// Read in config
	if err := parseConfig(); err != nil {

		// if no config file is found, create a default one
		var configErr viper.ConfigFileNotFoundError
		if errors.As(err, &configErr) {
			slog.Debug("Config file not found, creating default config")
			err := createDefaultConfig()
			return err
		}

		return err
	}

	return nil
}

// createDefaultConfig creates a default config file
func createDefaultConfig() error {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	// Default config path
	defaultConfigPath := path.Join(home, ".config", "sevp.toml")

	// Ensure the directory exists
	if err := os.MkdirAll(path.Dir(defaultConfigPath), 0750); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write the default config
	if err := os.WriteFile(defaultConfigPath, []byte(defaultConfig), 0600); err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}

	slog.Debug("Default config created", "path", defaultConfigPath)

	// Default file creation will still return an error
	// so the main CLI can exit once and prompt users to run sevp again
	return errors.New("created default config")
}

// ParseConfig reads in the config file.
//
// Viper looks for the config in $HOME/.config/sevp.toml and $HOME/sevp.toml.
// $Home/.config/sevp.toml takes precedence over $HOME/sevp.toml if both exist.
func parseConfig() error {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Set the config name and type
	viper.SetConfigName("sevp") // No dot here, Viper automatically adds extensions
	viper.SetConfigType("toml")

	viper.AddConfigPath(path.Join(home, ".config")) // $HOME/.config/sevp.toml

	// Read in the config
	if err := viper.ReadInConfig(); err != nil {
		slog.Debug("Error reading config", "err", err)
		return err
	}

	if len(viper.AllSettings()) == 0 {
		return errors.New("config file is empty")
	}

	slog.Debug("Config file read successfully", "path", viper.ConfigFileUsed())
	return nil
}

// GetSelector returns the appropriate selector based on CLI args and config.
func GetSelector(args []string) (Selector, error) {
	if viper.ConfigFileUsed() == "" {
		return nil, fmt.Errorf("no config file found")
	}

	// Check for default selector for when no args are provided
	selectorName := viper.GetString("default")

	// If target selector is provided, use it
	if len(args) == 1 {
		selectorName = args[0]
	}

	section, err := FromConfig(selectorName)
	if err != nil {
		return nil, err
	}

	// If the selector is an external config provider,
	// converts the config selector into a external config selector
	if section.ReadExternalConfig {
		return section.IntoExternalConfigSelector()
	}

	if section.TargetVar == "" || len(section.PossibleValues) == 0 {
		return nil, errors.New("missing target_var or possible_values")
	}

	return section, nil
}

// configSelectorMap maps selector name to ConfigSelector.
type ConfigSelectorMap map[string]*ConfigSelector

// ParseSelectorsFromConfig reads the config file and returns a map of defined selectors.
func ParseSelectorsFromConfig() (ConfigSelectorMap, error) {
	if viper.ConfigFileUsed() == "" {
		return nil, fmt.Errorf("no config file found")
	}

	selectors := make(ConfigSelectorMap)
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
