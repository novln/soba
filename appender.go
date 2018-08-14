package soba

import (
	"regexp"

	"github.com/pkg/errors"
)

const (
	// ConsoleAppenderType defines the type for a console appender.
	ConsoleAppenderType = "console"
	// FileAppenderType defines the type for a file appender.
	FileAppenderType = "file"
)

// An Appender defines an entity that receives a log entry and logs it somewhere,
// like for example, to a file, the console, or the syslog.
type Appender interface {
	// Name returns Appender name.
	Name() string
}

// IsAppenderNameValid verify that a Appender name has a valid format.
var IsAppenderNameValid = regexp.MustCompile(`^[a-z]+[a-z._0-9-]+[a-z0-9]+$`).MatchString

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

// Name returns Appender name.
func (appender *ConsoleAppender) Name() string {
	return appender.name
}

func (ConsoleAppender) flush() {

}

// TODO (novln): Add a rolling system to FileAppender

type FileAppender struct {
	name string
}

// Name returns Appender name.
func (appender *FileAppender) Name() string {
	return appender.name
}

func (FileAppender) flush() {

}
