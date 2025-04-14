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

		if !isValidTfenvVersionString(entry.Name()) {
			continue
		}

		versions = append(versions, entry.Name())
	}

	if len(versions) == 0 {
		return nil, errors.New("no tfenv versions")
	}

	return versions, nil
}

// isValidTfenvVersionString returns true if the version string is valid terraform version
func isValidTfenvVersionString(version string) bool {
	re := regexp.MustCompile(`^\d+\.\d+\.\d+(?:-(?:alpha\d+|beta\d+|rc\d+|oci|alpha\d{8}))?$`)
	return re.MatchString(version)
}
