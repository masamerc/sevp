package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
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

	// sanitize the file path
	filePath := filepath.Clean(filepath.Join(userHome, FileName))

	f, err := os.Create(filePath)
	if err != nil {
		slog.Debug("Error creating file", "path", filePath, "err", err)
		FailOnError("Error creating file", err)
	}
	defer f.Close()

	exportString := fmt.Sprintf("export %s=%s", target, value)
	_, err = f.WriteString(exportString + "\n")
	if err != nil {
		slog.Debug("Error writing to file", "path", filePath, "err", err)
		FailOnError("Error writing to file", err)
	}

	slog.Debug("Wrote environment variable to file", "path", filePath, "var", target, "value", value)
}
