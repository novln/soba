package soba

import (
	"context"
	"fmt"
	"regexp"
)

// A Logger provides fast, leveled, structured logging.
// All methods must be safe for concurrent use.
type Logger interface {
	// Debug logs a message at DebugLevel.
	Debug(message string, fields ...Field)
	// Info logs a message at InfoLevel.
	Info(message string, fields ...Field)
	// Warn logs a message at WarnLevel.
	Warn(message string, fields ...Field)
	// Error logs a message at ErrorLevel.
	Error(message string, fields ...Field)
	// With appends given structured fields to it.
	With(fields ...Field) Logger
}

// IsLoggerNameValid verify that a Logger name has a valid format.
var IsLoggerNameValid = regexp.MustCompile(`^[a-z]+[a-z._0-9]+[a-z0-9]+$`).MatchString

// New creates a new Logger using given name.
func New(ctx context.Context, name string) Logger {
	if !IsLoggerNameValid(name) {
		panic(fmt.Sprintf("soba: invalid logger name format: %s", name))
	}
	return nil
}
