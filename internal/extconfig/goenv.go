package extconfig

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

type GoEnvSelector struct{}

func (s GoEnvSelector) Read() (string, []string, error) {
	targetVar := "GOENV_VERSION"
	versions, err := getAvailableGoEnvVersions()
	return targetVar, versions, err
}

func NewGoEnvSelector() *GoEnvSelector {
	return &GoEnvSelector{}
}

func getAvailableGoEnvVersions() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return readGoEnvVersions(filepath.Join(home, ".goenv", "versions"))
}

func readGoEnvVersions(tfEnvPath string) ([]string, error) {
	entries, err := os.ReadDir(tfEnvPath)
	if err != nil {
		return nil, err
	}

	var versions []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if !isValidGoVersionString(entry.Name()) {
			continue
		}

		versions = append(versions, entry.Name())
	}

	if len(versions) == 0 {
		return nil, errors.New("no goenv versions")
	}

	return versions, nil
}

func isValidGoVersionString(version string) bool {
	re := regexp.MustCompile(`^\d+\.\d+(?:\.\d+)?(?:beta\d+|rc\d+)?$`)
	return re.MatchString(version)
}
