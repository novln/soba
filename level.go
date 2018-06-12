package soba

// Level define an entry priority.
type Level uint8

const (
	// UnknownLevel represents an unsupported level.
	UnknownLevel = Level(iota)
	// NoLevel is a no-op entry: the logger is disabled.
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
	strNoLevel         = "never"
	strNoLevelNo       = "no"
	strNoLevelNone     = "none"
	strNoLevelDisable  = "disable"
	strNoLevelDisabled = "disabled"
	strUnknownLevel    = "unknown"
	strVerboseLevel    = "verbose"
	strDebugLevel      = "debug"
	strInfoLevel       = "info"
	strWarnLevel       = "warning"
	strShortWarnLevel  = "warn"
	strErrorLevel      = "error"
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
func ParseLevel(level string) (Level, bool) {
	switch level {
	case strErrorLevel:
		return ErrorLevel, true
	case strWarnLevel, strShortWarnLevel:
		return WarnLevel, true
	case strInfoLevel:
		return InfoLevel, true
	case strDebugLevel, strVerboseLevel:
		return DebugLevel, true
	case strNoLevel, strNoLevelNo, strNoLevelNone, strNoLevelDisable, strNoLevelDisabled:
		return NoLevel, true
	default:
		return UnknownLevel, false
	}
}
