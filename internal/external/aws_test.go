package external

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Getting the AWS config file path should return the expected path
func TestGetConfigFile(t *testing.T) {
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	tempDir := t.TempDir()
	os.Setenv("HOME", tempDir)

	configPath, err := GetAWSConfigFile()

	assert.NoError(t, err, "expected no error getting config file")
	expectedPath := path.Join(tempDir, ".aws", "config")
	assert.Equal(t, expectedPath, configPath, "config path should match expected path")
}

// Reading the contents of a file should return the expected content without error
func TestReadContents(t *testing.T) {
	tempDir := t.TempDir()
	filePath := path.Join(tempDir, "testfile.txt")

	content := "test content"
	err := os.WriteFile(filePath, []byte(content), 0644)
	assert.NoError(t, err, "failed to create test file")

	result, err := readContents(filePath)
	assert.NoError(t, err, "expected no error reading file contents")
	assert.Equal(t, content, result, "file content should match")
}

// Parsing the AWS config file should return the expected profiles
func TestGetProfiles(t *testing.T) {
	tests := []struct {
		name     string
		contents string
		expected []string
	}{
		{
			name: "Single profile",
			contents: `
[default]
[profile my-profile]
`,
			expected: []string{"default", "my-profile"},
		},
		{
			name: "Multiple profiles",
			contents: `
[profile test1]
[profile test2]
[profile test3]
`,
			expected: []string{"test1", "test2", "test3"},
		},
		{
			name:     "No profiles",
			contents: "",
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseProfiles(test.contents)
			assert.Equal(t, test.expected, result, "profile list should match expected")
		})
	}
}
