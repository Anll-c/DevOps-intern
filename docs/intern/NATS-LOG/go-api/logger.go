package main

import (
	"log/slog"
	"os"
	"strings"
)

func Logger() {
	logLevel := strings.ToLower(os.Getenv("APP_LOG_LEVEL"))

	var level slog.Level
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})

	slog.SetDefault(slog.New(handler))

	slog.Info("Log started", "level", level.String())

}
