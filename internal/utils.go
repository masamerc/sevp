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

// FailOnError logs the error message and exits the program if an error is encountered.
//
// Parameters:
//   - msg: A descriptive message to log when an error occurs.
//   - err: The error to log and handle.
func FailOnError(msg string, err error) {
	if err != nil {
		slog.Error(err.Error())
		fmt.Println(msg)
		os.Exit(1)
	}
}

// InitLogger initializes the logger with the appropriate log level based on the SEVP_LOG_LEVEL
// environment variable.
//
// The supported log levels are:
//   - "debug": Debug-level logging.
//   - "info": Info-level logging (default).
//   - Any other value: Warning-level logging.
func InitLogger() {
	// log settings
	logLevelString := os.Getenv("SEVP_LOG_LEVEL")
	if logLevelString != "" {
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
//
// Parameters:
//   - value: The value of the environment variable.
//   - target: The name of the environment variable.
func WriteToFile(value string, target string) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		slog.Debug("Error getting user's home directory", "err", err)
		FailOnError("Error getting user's home directory", err)
	}

	filePath := filepath.Clean(filepath.Join(userHome, FileName))

	// Read existing file content
	var lines []string
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		slog.Debug("Error opening file", "path", filePath, "err", err)
		FailOnError("Error opening file", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close() // Close file after reading

	if err := scanner.Err(); err != nil {
		slog.Debug("Error scanning file", "path", filePath, "err", err)
		FailOnError("Error scanning file", err)
	}

	// Check if target exists and overwrite or append
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

	// Write updated content back to file
	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644) // Open for writing, truncate
	if err != nil {
		slog.Debug("Error opening file for writing", "path", filePath, "err", err)
		FailOnError("Error opening file for writing", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			slog.Debug("Error writing to file", "path", filePath, "err", err)
			FailOnError("Error writing to file", err)
		}
	}

	if err := writer.Flush(); err != nil {
		slog.Debug("Error flushing writer", "path", filePath, "err", err)
		FailOnError("Error flushing writer", err)
	}

	slog.Debug("Wrote environment variable to file", "path", filePath, "var", target, "value", value)
}
