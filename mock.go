package soba

import (
	"io"
)

// MockLogger is a mocked Logger which disable every operation.
type MockLogger struct {
}

// NewMockLogger returns a new MockLogger
func NewMockLogger() Logger {
	return &MockLogger{}
}

// Debug is a no-op for MockLogger.
func (MockLogger) Debug(message string, fields ...Field) {

}

// Info is a no-op for MockLogger.
func (MockLogger) Info(message string, fields ...Field) {

}

// Warn lis a no-op for MockLogger.
func (MockLogger) Warn(message string, fields ...Field) {

}

// Error is a no-op for MockLogger.
func (MockLogger) Error(message string, fields ...Field) {

}

// With is a no-op for MockLogger.
func (MockLogger) With(fields ...Field) Logger {
	return &MockLogger{}
}

// WithLevel is a no-op for MockLogger.
func (MockLogger) WithLevel(level Level) Logger {
	return &MockLogger{}
}

// WithOutput is a no-op for MockLogger.
func (MockLogger) WithOutput(writer io.Writer) Logger {
	return &MockLogger{}
}
