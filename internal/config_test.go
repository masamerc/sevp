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

//	func TestInvalidConfiguration(t *testing.T) {
//		tests := []struct {
//			name          string
//			configContent string
//			expectedErr   string
//		}{
//			{
//				name: "Missing Required Field TargetVar",
//				configContent: `
//
// [invalid]
// read_config = true
// possible_values = ["value1"]`,
//
//		expectedErr: "target_var",
//	},
//	{
//		name: "Missing Required Field PossibleValues",
//		configContent: `
//
// [invalid]
// read_config = true
// target_var = "TEST_VAR"`,
//
//		expectedErr: "possible_values",
//	},
//	{
//		name: "Invalid Type for ReadConfig",
//		configContent: `
//
// [invalid]
// read_config = "not_a_boolean"
// target_var = "TEST_VAR"
// possible_values = ["value1"]`,
//
//		expectedErr: "read_config",
//	},
//	{
//		name: "Invalid Type for PossibleValues",
//		configContent: `
//
// [invalid]
// read_config = true
// target_var = "TEST_VAR"
// possible_values = "not_an_array"`,
//
//				expectedErr: "possible_values",
//			},
//		}
//
//		for _, test := range tests {
//			t.Run(test.name, func(t *testing.T) {
//				viper.Reset()
//				viper.SetConfigType("toml")
//				err := viper.ReadConfig(strings.NewReader(test.configContent))
//				assert.NoError(t, err, "expected no error reading config")
//
//				_, err = FromConfig("invalid")
//				assert.Error(t, err, "expected error for invalid config")
//				assert.Contains(t, err.Error(), test.expectedErr, "error should mention the invalid field")
//			})
//		}
//	}

//	func TestEmptyConfiguration(t *testing.T) {
//		tests := []struct {
//			name          string
//			configContent string
//		}{
//			{
//				name:          "Empty Config File",
//				configContent: "",
//			},
//			{
//				name: "Empty Selector",
//				configContent: `
//
// [empty]`,
//
//			},
//		}
//
//		for _, test := range tests {
//			t.Run(test.name, func(t *testing.T) {
//				viper.Reset()
//				viper.SetConfigType("toml")
//				err := viper.ReadConfig(strings.NewReader(test.configContent))
//				assert.NoError(t, err, "expected no error reading config")
//
//				selectors, err := GetSelectors()
//				assert.NoError(t, err, "expected no error getting selectors")
//				assert.Empty(t, selectors, "selectors should be empty for empty config")
//			})
//		}
//	}
//
//	func TestEdgeCases(t *testing.T) {
//		configContent := `
//
// [empty_values]
// read_config = true
// target_var = "EMPTY_VAR"
// possible_values = []
//
// [single_value]
// read_config = true
// target_var = "SINGLE_VAR"
// possible_values = ["one"]
//
// [special_chars]
// read_config = true
// target_var = "SPECIAL_VAR!@#$"
// possible_values = ["value-with-dash", "value_with_underscore", "value with spaces", "!@#$%^&*()"]
//
// [unicode_chars]
// read_config = true
// target_var = "UNICODE_VAR"
// possible_values = ["值", "값", "🎉", "café"]
// `
//
//		viper.Reset()
//		viper.SetConfigType("toml")
//		err := viper.ReadConfig(strings.NewReader(configContent))
//		assert.NoError(t, err, "expected no error reading config")
//
//		tests := []struct {
//			name     string
//			selector string
//			check    func(*testing.T, *ConfigSelector)
//		}{
//			{
//				name:     "Empty Possible Values",
//				selector: "empty_values",
//				check: func(t *testing.T, s *ConfigSelector) {
//					assert.Empty(t, s.PossibleValues, "possible values should be empty")
//				},
//			},
//			{
//				name:     "Single Possible Value",
//				selector: "single_value",
//				check: func(t *testing.T, s *ConfigSelector) {
//					assert.Len(t, s.PossibleValues, 1, "should have exactly one possible value")
//					assert.Equal(t, "one", s.PossibleValues[0])
//				},
//			},
//			{
//				name:     "Special Characters in Values",
//				selector: "special_chars",
//				check: func(t *testing.T, s *ConfigSelector) {
//					assert.Len(t, s.PossibleValues, 4, "should have all special character values")
//					assert.Contains(t, s.PossibleValues, "value with spaces")
//					assert.Contains(t, s.PossibleValues, "!@#$%^&*()")
//				},
//			},
//			{
//				name:     "Unicode Characters",
//				selector: "unicode_chars",
//				check: func(t *testing.T, s *ConfigSelector) {
//					assert.Len(t, s.PossibleValues, 4, "should have all unicode values")
//					assert.Contains(t, s.PossibleValues, "値")
//					assert.Contains(t, s.PossibleValues, "🎉")
//				},
//			},
//		}
//
//		for _, test := range tests {
//			t.Run(test.name, func(t *testing.T) {
//				selector, err := FromConfig(test.selector)
//				assert.NoError(t, err, "expected no error creating selector")
//				test.check(t, selector)
//			})
//		}
//	}
func TestDuplicateSelectorsShouldNotParse(t *testing.T) {
	configContent := `
[duplicate]
read_config = true
target_var = "FIRST_VAR"
possible_values = ["first"]

[duplicate]
read_config = false
target_var = "SECOND_VAR"
possible_values = ["second"]
`
	viper.Reset()
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.Error(t, err, "expected error reading config with duplicate selectors")
	assert.Contains(t, err.Error(), "duplicate", "error should mention the duplicate table")
}

func TestNonExistentSelector(t *testing.T) {
	configContent := `
[existing]
read_config = true
target_var = "TEST_VAR"
possible_values = ["value"]
`
	viper.Reset()
	viper.SetConfigType("toml")
	err := viper.ReadConfig(strings.NewReader(configContent))
	assert.NoError(t, err, "expected no error reading config")

	_, err = FromConfig("non_existent")
	assert.Error(t, err, "expected error for non-existent selector")
	assert.Contains(t, err.Error(), "the selector is not in the config", "error should mention the missing selector")
}

func TestNoConfigFile(t *testing.T) {
	_, err := FromConfig("no config")
	assert.Error(t, err, "expected error")
	assert.Contains(t, err.Error(), "the config file is not found", "error should mention the missing config file")

	_, err = GetSelectors()
	assert.Error(t, err, "expected error")
	assert.Contains(t, err.Error(), "no config file found", "error should mention the missing config file")
}
