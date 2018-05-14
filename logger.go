package soba

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
	// With clones the current Logger and append given structured context to it.
	With(fields ...Field) Logger
}
