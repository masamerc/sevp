package internal

import (
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// AWSProfileSelector is a struct that implements the Selector interface for selecting AWS profiles.
type AWSProfileSelector struct{}

// Read is the main function of the AWSProfileSelector struct, which reads the AWS profile names from the AWS config file.
func (s *AWSProfileSelector) Read() (string, []string, error) {
	targetVar := "AWS_PROFILE"
	profiles, err := getAWSProfiles()
	return targetVar, profiles, err
}

// NewAWSProfileSelector creates a new empty instance of AWSProfileSelector.
func NewAWSProfileSelector() *AWSProfileSelector {
	return &AWSProfileSelector{}
}

// GetAWSConfigFile retrieves the path to the AWS config file.
//
// This operation can fail if the home directory cannot be determined for some reason.
func GetAWSConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Debug("Error getting home directory", "err", err)
		return "", err
	}

	slog.Debug("Read config file", "home", home)
	awsConfigPath := path.Join(home, ".aws", "config")
	return awsConfigPath, nil
}

// readContents reads the contents of a file given its path.
//
// This operation can fail if reading the file fails or if the file does not exist.
func readContents(filePath string) (string, error) {
	// sanitize the file path
	filePath = filepath.Clean(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		slog.Debug("Error opening file", "path", filePath, "err", err)
		return "", err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		slog.Debug("Error reading file", "path", filePath, "err", err)
		return "", err
	}

	return string(buf), nil
}

// parseProfiles extracts AWS profile names from the AWS config file contents.
//
// In case of no matches, it just returns an empty list and not throws an error.
func parseProfiles(contents string) []string {
	pattern := regexp.MustCompile(`\[(?:profile)?(.*?)\]`)
	matches := pattern.FindAllStringSubmatch(contents, -1)

	if matches == nil {
		return []string{}
	}

	var result []string
	for _, match := range matches {
		matched := strings.TrimSpace(match[1])
		result = append(result, matched)
	}

	return result
}

// getAWSProfiles retrieves a list of AWS profile names from the user's AWS config file.
//
// If it fails to either get the config file or read its contents, it returns an empty list and an error.
func getAWSProfiles() ([]string, error) {
	configPath, err := GetAWSConfigFile()
	if err != nil {
		return nil, err
	}

	contents, err := readContents(configPath)
	if err != nil {
		return nil, err
	}

	return parseProfiles(contents), nil
}
