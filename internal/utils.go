package internal

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const (
	FileName = ".sevp"
)

// InitLogger initializes the logger with the appropriate log level based on the SEVP_LOG_LEVEL.
//
// The log level can be set to "debug", "info", or "warn". If not set, it defaults to "info".
func InitLogger() {
	// log settings
	logLevelString := os.Getenv("SEVP_LOG_LEVEL")
	if logLevelString == "" {
		logLevelString = "info"
	}
	switch logLevelString {
	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelWarn)
	}
}

// WriteToFile writes an environment variable to a file.
func WriteToFile(value string, target string) error {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filePath := filepath.Clean(filepath.Join(userHome, FileName))

	// read existing file content
	var lines []string
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := file.Close(); err != nil {
		return err
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// check if target exists and overwrite or append
	targetFound := false
	for i, line := range lines {
		if strings.HasPrefix(line, fmt.Sprintf("export %s=", target)) {
			lines[i] = fmt.Sprintf("export %s=%s", target, value)
			targetFound = true
			break
		}
	}

	if !targetFound {
		lines = append(lines, fmt.Sprintf("export %s=%s", target, value))
	}

	// write updated content back to file
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600) // open for writing, truncate
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	slog.Debug("Wrote environment variable to file", "path", filePath, "var", target, "value", value)

	return nil
}
