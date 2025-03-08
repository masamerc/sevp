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

type AWSProfileSelector struct{}

func (s *AWSProfileSelector) Read() (string, []string, error) {
	targetVar := "AWS_PROFILE"
	profiles, err := getAWSProfiles()
	return targetVar, profiles, err
}

func NewAWSProfileSelector() *AWSProfileSelector {
	return &AWSProfileSelector{}
}

// GetConfigFile retrieves the path to the AWS config file.
//
// Returns:
//   - string: The full path to the AWS config file.
//   - error: An error if the user's home directory cannot be determined.
func getAWSConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Debug("Error getting home directory", "err", err)
		return "", err
	}

	slog.Debug("Read config file", "home", home)
	awsConfigPath := path.Join(home, ".aws", "config")
	return awsConfigPath, nil
}

// ReadContents reads the contents of a file given its path.
//
// Parameters:
//   - filePath: The full path to the file to read.
//
// Returns:
//   - string: The file's contents as a string.
//   - error: An error if the file cannot be opened or read.
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

// GetProfiles extracts AWS profile names from the AWS config file contents.
//
// Parameters:
//   - contents: The contents of the AWS config file as a string.
//
// Returns:
//   - []string: A list of profile names found in the config file.
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

// GetAWSProfiles retrieves a list of AWS profile names from the user's AWS config file.
//
// Returns:
//   - []string: A list of AWS profile names.
//   - error: An error if the AWS config file cannot be read.
func getAWSProfiles() ([]string, error) {
	configPath, err := getAWSConfigFile()
	if err != nil {
		return nil, err
	}

	contents, err := readContents(configPath)
	if err != nil {
		return nil, err
	}

	return parseProfiles(contents), nil
}
