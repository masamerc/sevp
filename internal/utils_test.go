package internal

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Writing to a file should create the file if it doesn't exist
func TestWriteToFile(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, ".sevp")

	// set the user's home directory to the temporary directory
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempDir)

	// test writing a new environment variable
	err := WriteToFile("test_value", "TEST_VAR")
	assert.NoError(t, err, "expected no error writing to file")

	// verify the file
	content, err := os.ReadFile(tempFile)
	assert.NoError(t, err, "expected no error reading file")
	assert.Contains(t, string(content), "export TEST_VAR=test_value", "file content should contain the environment variable")

	// test updating an existing environment variable
	err = WriteToFile("new_value", "TEST_VAR")
	assert.NoError(t, err, "expected no error writing to file")

	// verify the updated file content
	content, err = os.ReadFile(tempFile)
	assert.NoError(t, err, "expected no error reading file")
	assert.Contains(t, string(content), "export TEST_VAR=new_value", "file content should contain the updated environment variable")

	// test adding another environment variable
	err = WriteToFile("another_value", "ANOTHER_VAR")
	assert.NoError(t, err, "expected no error writing to file")

	// verify the file content
	content, err = os.ReadFile(tempFile)
	assert.NoError(t, err, "expected no error reading file")
	assert.Contains(t, string(content), "export ANOTHER_VAR=another_value", "file content should contain the new environment variable")

	// verify the resulting file contains all environment variables
	content, err = os.ReadFile(tempFile)
	assert.NoError(t, err, "expected no error reading file")
	assert.Contains(t, string(content), "export TEST_VAR=new_value", "file content should contain the updated environment variable")
	assert.Contains(t, string(content), "export ANOTHER_VAR=another_value", "file content should contain the new environment variable")
}

// Initializing the logger should set the log level based on the SEVP_LOG_LEVEL environment variable
func TestInitLogger(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		envValue    string
		expectedLog slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"", slog.LevelInfo}, // default case
		{"other", slog.LevelWarn},
	}

	for _, test := range tests {
		t.Run(test.envValue, func(t *testing.T) {
			os.Setenv("SEVP_LOG_LEVEL", test.envValue)
			InitLogger()
			assert.True(t, slog.Default().Enabled(ctx, test.expectedLog), "log level mismatch")
		})
	}
}
