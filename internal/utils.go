package internal

import (
	"fmt"
	"log/slog"
	"os"
)

func FailOnError(msg string, err error) {
	if err != nil {
		slog.Error(err.Error())
		fmt.Println(msg)
		os.Exit(1)
	}
}

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
