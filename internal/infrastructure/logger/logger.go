package logger

import (
	"log/slog"
	"os"
	"strings"
)

// New creates a new structured logger
func New(level, environment string) *slog.Logger {
	var logLevel slog.Level
	
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn", "warning":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	var handler slog.Handler
	if environment == "production" {
		// JSON format for production
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		// Text format for development
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
