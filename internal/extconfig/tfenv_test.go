package extconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestReadTfenvVersions should return the names of all tfenv versions
func TestReadTfenvVersions(t *testing.T) {
	tmp := t.TempDir()

	// Simulate tfenv version structure
	version1 := filepath.Join(tmp, "0.1.0")
	version2 := filepath.Join(tmp, "11.2.0")
	_ = os.MkdirAll(version1, 0750)
	_ = os.MkdirAll(version2, 0750)

	v, err := readTfenvVersions(tmp)

	require.NoError(t, err)
	require.ElementsMatch(t, []string{"0.1.0", "11.2.0"}, v)
}

// TestReadTfenvVersionsInvalidVersion should ignore invalid versions
func TestReadTfenvVersionsInvalidVersion(t *testing.T) {
	tmp := t.TempDir()

	// Simulate tfenv version structure
	version1 := filepath.Join(tmp, "0.1.0.100") // Invalid version
	version2 := filepath.Join(tmp, "11.2.0")
	_ = os.MkdirAll(version1, 0750)
	_ = os.MkdirAll(version2, 0750)

	v, err := readTfenvVersions(tmp)

	require.NoError(t, err)

	// The invalid version should be ignored
	require.ElementsMatch(t, []string{"11.2.0"}, v)
}

// TestReadTfenvVersionsEmpty should return an error if the tfenv dir is empty
func TestReadTfenvVersionsEmpty(t *testing.T) {
	tmp := t.TempDir()
	_, err := readTfenvVersions(tmp)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no tfenv versions")
}
