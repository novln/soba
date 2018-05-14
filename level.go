package soba

import (
	"github.com/pkg/errors"
)

// Level define an entry priority.
type Level uint8

const (
	// UnknownLevel represents an unsupported level.
	UnknownLevel = Level(iota)
	// NoLevel is a no-op entries: the logger is disabled.
	NoLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// no error should be generated.
	ErrorLevel
	// WarnLevel is a non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel is the default logging priority: general operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging
	DebugLevel
)

const (
	strNoLevel        = "never"
	strNoLevelNo      = "no"
	strNoLevelNone    = "none"
	strUnknownLevel   = "unknown"
	strVerboseLevel   = "verbose"
	strDebugLevel     = "debug"
	strInfoLevel      = "info"
	strWarnLevel      = "warning"
	strShortWarnLevel = "warn"
	strErrorLevel     = "error"
)

// Convert the Level to a string.
func (level Level) String() string {
	switch level {
	case DebugLevel:
		return strDebugLevel
	case InfoLevel:
		return strInfoLevel
	case WarnLevel:
		return strWarnLevel
	case ErrorLevel:
		return strErrorLevel
	case NoLevel:
		return strNoLevel
	default:
		return strUnknownLevel
	}
}

// ParseLevel takes a string level and returns the log level constant.
func ParseLevel(level string) (Level, error) {
	switch level {
	case strErrorLevel:
		return ErrorLevel, nil
	case strWarnLevel, strShortWarnLevel:
		return WarnLevel, nil
	case strInfoLevel:
		return InfoLevel, nil
	case strDebugLevel, strVerboseLevel:
		return DebugLevel, nil
	case strNoLevel, strNoLevelNo, strNoLevelNone:
		return NoLevel, nil
	default:
		return UnknownLevel, errors.Errorf("not a valid logger Level: %q", level)
	}
}
