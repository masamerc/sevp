package internal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
