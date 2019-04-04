package soba

import (
	"context"
	"fmt"
	"regexp"
)

// IsLoggerNameValid verify that a Logger name has a valid format.
var IsLoggerNameValid = regexp.MustCompile(`^[a-z]+[a-z._0-9-]+[a-z0-9]+$`).MatchString

// A Logger provides fast, leveled, structured logging.
// All methods are safe for concurrent use.
type Logger struct {
	name      string
	level     Level
	appenders []Appender
	fields    []Field
}

// New creates a new Logger using given name.
func New(ctx context.Context, name string) Logger {
	if !IsLoggerNameValid(name) {
		panic(fmt.Sprintf("soba: invalid logger name format: %s", name))
	}

	handler := ctx.Value(hCtxKey).(Handler)
	if handler == nil {
		panic("soba: must be initialized with soba.Load()")
	}

	return handler.New(name)
}

// NewLogger creates a new logger.
func NewLogger(name string, level Level, appenders []Appender) Logger {
	if !IsLoggerNameValid(name) {
		panic(fmt.Sprintf("soba: invalid logger name format: %s", name))
	}
	return Logger{
		name:      name,
		level:     level,
		appenders: appenders,
		fields:    make([]Field, 0, 64),
	}
}

// Name returns logger name.
func (logger Logger) Name() string {
	return logger.name
}

// Level returns logger level.
func (logger Logger) Level() Level {
	return logger.level
}

// Debug logs a message at DebugLevel.
func (logger Logger) Debug(message string, fields ...Field) {
	if logger.level < DebugLevel || logger.level == NoLevel {
		return
	}
	logger.write(DebugLevel, message, fields)
}

// Info logs a message at InfoLevel.
func (logger Logger) Info(message string, fields ...Field) {
	if logger.level < InfoLevel || logger.level == NoLevel {
		return
	}
	logger.write(InfoLevel, message, fields)
}

// Warn logs a message at WarnLevel.
func (logger Logger) Warn(message string, fields ...Field) {
	if logger.level < WarnLevel || logger.level == NoLevel {
		return
	}
	logger.write(WarnLevel, message, fields)
}

// Error logs a message at ErrorLevel.
func (logger Logger) Error(message string, fields ...Field) {
	if logger.level < ErrorLevel || logger.level == NoLevel {
		return
	}
	logger.write(ErrorLevel, message, fields)
}

// With appends given structured fields to it.
func (logger Logger) With(fields ...Field) Logger {
	other := logger.copy()
	other.fields = append(other.fields, fields...)
	return other
}

func (logger Logger) write(level Level, message string, fields []Field) {
	entry := NewEntry(logger.name, level, message, logger.fields, fields)
	defer entry.Flush()
	for i := range logger.appenders {
		logger.appenders[i].Write(entry)
	}
}

func (logger Logger) copyWithName(name string) Logger {
	if !IsLoggerNameValid(name) {
		panic(fmt.Sprintf("soba: invalid logger name format: %s", name))
	}

	other := logger.copy()
	other.name = name

	return other
}

func (logger Logger) copy() Logger {
	other := Logger{}

	other.name = logger.name
	other.level = logger.level
	other.appenders = logger.appenders

	other.fields = make([]Field, len(logger.fields), cap(logger.fields))
	copy(other.fields, logger.fields)

	return other
}
