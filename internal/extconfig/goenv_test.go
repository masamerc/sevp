package extconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestReadGoEnvVersions should return the names of all Go versions
func TestReadGoEnvVersions(t *testing.T) {
	tmp := t.TempDir()

	// Simulate goenv version structure
	version1 := filepath.Join(tmp, "1.18.0")
	version2 := filepath.Join(tmp, "1.19.1")
	_ = os.MkdirAll(version1, 0750)
	_ = os.MkdirAll(version2, 0750)

	versions, err := readGoEnvVersions(tmp)

	require.NoError(t, err)
	require.ElementsMatch(t, []string{"1.18.0", "1.19.1"}, versions)
}

// TestReadGoEnvVersionsInvalidVersion should ignore invalid versions
func TestReadGoEnvVersionsInvalidVersion(t *testing.T) {
	tmp := t.TempDir()

	// Simulate goenv version structure
	version1 := filepath.Join(tmp, "1.18.0.100")   // Invalid version
	version2 := filepath.Join(tmp, "1.19.1")       // Valid version
	version3 := filepath.Join(tmp, "1.20beta1")    // Valid version
	version4 := filepath.Join(tmp, "1.20.0-beta1") // Invalid version
	_ = os.MkdirAll(version1, 0750)
	_ = os.MkdirAll(version2, 0750)
	_ = os.MkdirAll(version3, 0750)
	_ = os.MkdirAll(version4, 0750)

	versions, err := readGoEnvVersions(tmp)

	require.NoError(t, err)

	// The invalid version should be ignored
	require.ElementsMatch(t, []string{"1.19.1", "1.20beta1"}, versions)
}

// TestReadGoEnvVersionsEmpty should return an error if the goenv dir is empty
func TestReadGoEnvVersionsEmpty(t *testing.T) {
	tmp := t.TempDir()
	_, err := readGoEnvVersions(tmp)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no goenv versions")
}
