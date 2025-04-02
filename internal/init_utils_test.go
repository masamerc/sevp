package internal

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	// backup and restore environment variables
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// set up a temporary home directory
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)

	// test case: no config file and no aws config
	err := InitConfig()
	assert.Error(t, err, "Expected error when no config files are present")

	// test case: valid aws config file
	awsConfigPath := tempHome + "/.aws/config"
	os.MkdirAll(tempHome+"/.aws", 0755)
	os.WriteFile(awsConfigPath, []byte("[default]\nregion=us-east-1"), 0644)

	err = InitConfig()
	assert.NoError(t, err, "Expected no error when AWS config file is present")
}

func TestInitSelector(t *testing.T) {
	// backup and restore viper configuration
	originalConfig := viper.AllSettings()
	defer func() {
		viper.Reset()
		for key, value := range originalConfig {
			viper.Set(key, value)
		}
	}()

	// test case: no config file used
	selector, err := InitSelector([]string{})
	assert.NoError(t, err, "Expected no error when no config file is used")
	assert.NotNil(t, selector, "Expected a valid selector")

	// test case: invalid selector in config
	viper.SetConfigFile("non-existent.toml")
	viper.Set("invalid_selector.read_config", false)
	viper.Set("invalid_selector.target_var", "")
	viper.Set("invalid_selector.possible_values", []string{})

	s, err := InitSelector([]string{"invalid_selector"})
	t.Log("Selector:", s)
	t.Log("Error:", err)
	assert.Error(t, err, "Expected error for invalid selector configuration")

	// test case: valid selector in config
	viper.Set("valid_selector.read_config", false)
	viper.Set("valid_selector.target_var", "TEST_VAR")
	viper.Set("valid_selector.possible_values", []string{"value1", "value2"})

	selector, err = InitSelector([]string{"valid_selector"})
	assert.NoError(t, err, "Expected no error for valid selector configuration")
	assert.NotNil(t, selector, "Expected a valid selector")
}
