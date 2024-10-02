package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"regexp"
	"strings"
)

// read from $HOME/.aws/config

func GetConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Debug("Error getting home directory", "err", err)
		return "", err
	}

	slog.Debug("Read config file", "home", home)

	awsConfigPath := path.Join(home, ".aws", "config")

	return awsConfigPath, nil
}

func ReadContents(awsConfigPath string) (string, error) {
	// read contents
	file, err := os.Open(awsConfigPath)
	if err != nil {
		slog.Debug("Error opening file", "err", err)
		return "", err
	}

	buf, err := io.ReadAll(file)
	if err != nil {
		slog.Debug("Error reading file", "err", err)
		return "", err
	}

	// []byte to string
	contents := string(buf)

	return contents, nil
}

func GetProfiles(contents string) []string {
	// extract [profile_name] [profile profile_name] using regex
	pattern := regexp.MustCompile(`\[(?:profile)?(.*?)\]`)
	matches := pattern.FindAllStringSubmatch(contents, -1)

	var result []string

	for _, match := range matches {
		matched := match[1]
		matched = strings.TrimSpace(matched)
		result = append(result, matched)
	}

	return result
}

func WriteToFile(profile string) {
	// create file ~/.sevp
	userHome, err := os.UserHomeDir()
	FailOnError("Error getting user's home directory", err)

	f, err := os.Create(path.Join(userHome, ".sevp"))
	FailOnError("Error creating file ~/.sevp", err)

	defer f.Close()

	// write profile to file
	exportString := fmt.Sprintf("export AWS_PROFILE=%s", profile)
	_, err = fmt.Fprintln(f, exportString)
	FailOnError("Error writing to file ~/.sevp", err)
	slog.Debug("Wrote profile to file ~/.sevp", "AWS_PROFILE", profile)
}
