package internal

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

type selectorMap map[string]*SelectorSection

type SelectorSection struct {
	Name           string
	ReadConfig     bool
	TargetVar      string
	PossibleValues []string
}

func fromConfig(name string) (*SelectorSection, error) {
	readConfig := viper.GetBool(name + ".read_config")
	targetVar := viper.GetString(name + ".target_var")
	possibleValues := viper.GetStringSlice(name + ".possible_values")

	if targetVar == "" || len(possibleValues) == 0 {
		return nil, fmt.Errorf("missing target_var or possible_values")
	}

	return &SelectorSection{
		Name:           name,
		ReadConfig:     readConfig,
		TargetVar:      targetVar,
		PossibleValues: possibleValues,
	}, nil
}

func ParseConfig() {
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
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
}

func GetSelectors() selectorMap {
	selectors := make(selectorMap)
	for key := range viper.AllSettings() {
		if key == "default" {
			continue
		}
		s, err := fromConfig(key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		selectors[key] = s
	}

	return selectors
}
