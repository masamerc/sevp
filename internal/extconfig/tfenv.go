package extconfig

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

type TfEnvSelector struct{}

func (s TfEnvSelector) Read() (string, []string, error) {
	targetVar := "TFENV_TERRAFORM_VERSION"
	versions, err := getAvailableTfenvVersions()
	return targetVar, versions, err
}

func NewTfEnvSelector() *TfEnvSelector {
	return &TfEnvSelector{}
}

// getAvailableTfenvVersions returns all available versions of Terraform managed by tfenv
func getAvailableTfenvVersions() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return readTfenvVersions(filepath.Join(home, ".tfenv", "versions"))
}

// readTfEnvVersions returns all available versions of Terraform managed by tfenv
func readTfenvVersions(tfEnvPath string) ([]string, error) {
	entries, err := os.ReadDir(tfEnvPath)
	if err != nil {
		return nil, err
	}

	var versions []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if !isValidVersionString(entry.Name()) {
			continue
		}

		versions = append(versions, entry.Name())
	}

	if len(versions) == 0 {
		return nil, errors.New("no tfenv versions")
	}

	return versions, nil
}

// isValidVersionString checks if the given string is a valid Terraform version string (e.g. 1.1.0)
func isValidVersionString(versionString string) bool {
	pattern := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return pattern.MatchString(versionString)
}
