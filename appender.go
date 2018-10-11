package soba

import (
	"fmt"
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
	// Write receives a log entry.
	Write(entry Entry)
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

// ConsoleAppender is an appender that uses stdout to write log entry.
type ConsoleAppender struct {
	name string
}

// Name returns Appender name.
func (appender *ConsoleAppender) Name() string {
	return appender.name
}

// Write receives a log entry.
func (appender *ConsoleAppender) Write(entry Entry) {
	// TODO Optimize this ?
	fmt.Println(entry)
}

// TODO (novln): Add a rolling system to FileAppender

// FileAppender is an appender that uses a file to write log entry.
type FileAppender struct {
	name string
}

// Name returns Appender name.
func (appender *FileAppender) Name() string {
	return appender.name
}

// Write receives a log entry.
func (appender *FileAppender) Write(entry Entry) {
	// TODO
	fmt.Println(entry)
}
