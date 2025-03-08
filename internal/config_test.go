package internal

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestFromConfig(t *testing.T) {
	configContent := `
[aws]
read_config = true
target_var = "AWS_PROFILE"
possible_values = ["default", "profile1", "profile2"]

[custom]
read_config = false
target_var = "CUSTOM_VAR"
possible_values = ["value1", "value2"]
`
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.NoError(t, err, "expected no error reading config")

	tests := []struct {
		name     string
		selector string
		expected *ConfigSelector
	}{
		{
			name:     "AWS Selector",
			selector: "aws",
			expected: &ConfigSelector{
				Name:           "aws",
				ReadConfig:     true,
				TargetVar:      "AWS_PROFILE",
				PossibleValues: []string{"default", "profile1", "profile2"},
			},
		},
		{
			name:     "Custom Selector",
			selector: "custom",
			expected: &ConfigSelector{
				Name:           "custom",
				ReadConfig:     false,
				TargetVar:      "CUSTOM_VAR",
				PossibleValues: []string{"value1", "value2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := FromConfig(test.selector)
			assert.NoError(t, err, "expected no error creating selector")
			assert.Equal(t, test.expected, result, "selector should match expected")
		})
	}
}

func TestParseConfig(t *testing.T) {
	configContent := `
[aws]
read_config = true
target_var = "AWS_PROFILE"
possible_values = ["default", "profile1", "profile2"]
`
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.NoError(t, err, "expected no error reading config")

	err = ParseConfig()
	assert.NoError(t, err, "expected no error parsing config")
}

func TestGetSelectors(t *testing.T) {
	configContent := `
[aws]
read_config = true
target_var = "AWS_PROFILE"
possible_values = ["default", "profile1", "profile2"]

[custom]
read_config = false
target_var = "CUSTOM_VAR"
possible_values = ["value1", "value2"]
`
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.NoError(t, err, "expected no error reading config")

	selectors, err := GetSelectors()
	assert.NoError(t, err, "expected no error getting selectors")

	expectedSelectors := configSelectorMap{
		"aws": &ConfigSelector{
			Name:           "aws",
			ReadConfig:     true,
			TargetVar:      "AWS_PROFILE",
			PossibleValues: []string{"default", "profile1", "profile2"},
		},
		"custom": &ConfigSelector{
			Name:           "custom",
			ReadConfig:     false,
			TargetVar:      "CUSTOM_VAR",
			PossibleValues: []string{"value1", "value2"},
		},
	}

	assert.Equal(t, expectedSelectors, selectors, "selectors should match expected")
}

func TestDefaultSelector(t *testing.T) {
	configContent := `
default = "aws"

[aws]
read_config = true
target_var = "AWS_PROFILE"
possible_values = ["default", "profile1", "profile2"]
`
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.NoError(t, err, "expected no error reading config")

	selectorChoice := viper.GetString("default")

	selectors, err := GetSelectors()
	assert.NoError(t, err, "expected no error getting selectors")

	selector, ok := selectors[selectorChoice]
	assert.True(t, ok, "expected selector to be found")

	expectedSelector := &ConfigSelector{
		Name:           "aws",
		ReadConfig:     true,
		TargetVar:      "AWS_PROFILE",
		PossibleValues: []string{"default", "profile1", "profile2"},
	}

	assert.Equal(t, expectedSelector, selector, "selectors should match expected")
}
