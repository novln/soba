package soba

import (
	"github.com/pkg/errors"
)

const (
	// ConsoleAppenderType defines the type for a console appender.
	ConsoleAppenderType = "console"
	// FileAppenderType defines the type for a file appender.
	FileAppenderType = "file"
)

type Appender interface {
	// Name returns Appender name.
	Name() string
	flush()
}

// NewAppender creates a new Appender from given configuration.
func NewAppender(name string, conf ConfigAppender) (Appender, error) {
	switch conf.Type {
	case ConsoleAppenderType:
		appender := &ConsoleAppender{
			name: name,
		}

		return appender, nil

	case FileAppenderType:
		appender := &FileAppender{
			name: name,
		}

		return appender, nil

	default:
		return nil, errors.Errorf("unknown appender type: %s", conf.Type)
	}
}

type ConsoleAppender struct {
	name string
}

func (appender *ConsoleAppender) Name() string {
	return appender.name
}

func (ConsoleAppender) flush() {

}

// TODO (novln): Add a rolling system to FileAppender

type FileAppender struct {
	name string
}

func (appender *FileAppender) Name() string {
	return appender.name
}

func (FileAppender) flush() {

}
