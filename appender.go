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
	flush()
}

// NewAppender creates a new Appender from given configuration.
func NewAppender(conf ConfigAppender) (Appender, error) {
	switch conf.Type {
	case ConsoleAppenderType:
		return ConsoleAppender{}, nil

	case FileAppenderType:
		return FileAppender{}, nil

	default:
		return nil, errors.Errorf("unknown appender type: %s", conf.Type)
	}
}

type ConsoleAppender struct {
}

func (ConsoleAppender) flush() {

}

// TODO (novln): Add a rolling system to FileAppender

type FileAppender struct {
}

func (FileAppender) flush() {

}
