package soba

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"

	"github.com/pkg/errors"

	"github.com/novln/soba/encoder/json"
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
	// Name returns appender name.
	Name() string
	// Write receives a log entry.
	Write(entry *Entry)
	// Close recycles underlying resources of appender.
	Close() error
}

// IsAppenderNameValid verify that a Appender name has a valid format.
var IsAppenderNameValid = regexp.MustCompile(`^[a-z]+[a-z._0-9-]+[a-z0-9]+$`).MatchString

// NewAppender creates a new Appender from given configuration.
// To register a custom appender, please use soba.RegisterAppenders() function.
func NewAppender(name string, conf ConfigAppender) (Appender, error) {
	err := validateAppenderConfig(name, conf)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create appender for %s", name)
	}

	switch conf.Type {
	case ConsoleAppenderType:
		appender := NewConsoleAppender(name, os.Stdout)
		return appender, nil

	case FileAppenderType:
		appender, err := NewFileAppender(name, conf.Path, conf.Backup, conf.MaxBytes)
		if err != nil {
			return nil, errors.Wrapf(err, "cannot create file appender for %s", name)
		}

		return appender, nil

	default:
		// Should be handled by validateAppenderConfig function.
		return nil, errors.Errorf("unknown appender type for %s: %s", name, conf.Type)
	}
}

// ConsoleAppender is an appender that uses stdout to write log entry.
type ConsoleAppender struct {
	mutex sync.Mutex
	name  string
	out   io.Writer
}

// NewConsoleAppender creates a new ConsoleAppender instance.
func NewConsoleAppender(name string, out io.Writer) *ConsoleAppender {
	return &ConsoleAppender{
		name: name,
		out:  out,
	}
}

// Name returns appender name.
func (appender *ConsoleAppender) Name() string {
	return appender.name
}

// Close recycles underlying resources of appender.
func (appender *ConsoleAppender) Close() error {
	return nil
}

// Write receives a log entry.
func (appender *ConsoleAppender) Write(entry *Entry) {
	encoder := json.NewEncoder()
	defer encoder.Close()

	buffer := WriteEntry(entry, encoder)

	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	_, err := appender.out.Write(buffer)
	if err != nil {
		onAppenderWriteError(err)
	}
}

// FileAppender is an appender that uses a file to write log entry.
type FileAppender struct {
	mutex    sync.Mutex
	name     string
	path     string
	file     *os.File
	size     int64
	backup   bool
	maxBytes int64
}

// NewFileAppender creates a new FileAppender instance.
func NewFileAppender(name string, path string, backup bool, maxBytes int64) (*FileAppender, error) {
	appender := &FileAppender{
		name:     name,
		path:     path,
		backup:   backup,
		maxBytes: maxBytes,
	}

	err := appender.openNew()
	if err != nil {
		return nil, errors.Wrapf(err, `cannot create file "%s" for appender`, path)
	}

	return appender, nil
}

// Name returns appender name.
func (appender *FileAppender) Name() string {
	return appender.name
}

// Close recycles underlying resources of appender.
func (appender *FileAppender) Close() error {
	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	return appender.close()
}

// Write receives a log entry and writes it on a file.
func (appender *FileAppender) Write(entry *Entry) {
	encoder := json.NewEncoder()
	defer encoder.Close()

	buffer := WriteEntry(entry, encoder)

	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	err := appender.rotate(len(buffer))
	if err != nil {
		onAppenderWriteError(err)
		return
	}

	n, err := appender.file.Write(buffer)
	if err != nil {
		onAppenderWriteError(err)
		return
	}

	appender.size += int64(n)
}

// openNew opens a new appender file to write log entry.
func (appender *FileAppender) openNew() error {
	if appender.file != nil {
		return nil
	}

	directory := filepath.Dir(appender.path)
	err := os.MkdirAll(directory, 0750)
	if err != nil {
		return errors.WithStack(err)
	}

	file, err := os.OpenFile(appender.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return errors.WithStack(err)
	}

	fstat, err := file.Stat()
	if err != nil {
		return errors.WithStack(err)
	}

	appender.file = file
	appender.size = fstat.Size()

	return nil
}

// close closes the underlying file.
func (appender *FileAppender) close() error {
	if appender.file == nil {
		return nil
	}

	file := appender.file
	appender.file = nil
	appender.size = 0

	err := file.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// rotate analyzes if a file rotation is required, and executes it when needed.
func (appender *FileAppender) rotate(toWrite int) error {
	if appender.maxBytes == 0 {
		return nil
	}

	finalSize := int64(toWrite) + appender.size
	if finalSize < appender.maxBytes {
		return nil
	}

	err := appender.close()
	if err != nil {
		return err
	}

	err = appender.doRotate()
	if err != nil {
		return err
	}

	err = appender.openNew()
	if err != nil {
		return err
	}

	return nil
}

// doRotate executes the file rotation, using backup or rename strategy.
func (appender *FileAppender) doRotate() error {
	if appender.backup {
		return appender.rotateWithBackup()
	}
	return appender.rotateWithRename()
}

// rotateWithRename renames the current file by suffixing it with "-".
// If a file with the suffix already exists, it will be replaced.
func (appender *FileAppender) rotateWithRename() error {
	backup := fmt.Sprint(appender.path, "-")

	err := os.Rename(appender.path, backup)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// rotateWithBackup renames the current file by suffixing it with a backup number.
//
// Let's say with we the following files in the directory:
//   - app.log
//   - app.log.1
//   - app.log.2
//   - app.log.3
//
// The current file (app.log) will be rename to "app.log.4".
func (appender *FileAppender) rotateWithBackup() error {
	pattern := fmt.Sprint(appender.path, ".*")
	list, err := filepath.Glob(pattern)
	if err != nil {
		return errors.WithStack(err)
	}

	id := len(list) + 1
	backup := fmt.Sprint(appender.path, ".", strconv.FormatInt(int64(id), 10))

	err = os.Rename(appender.path, backup)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func onAppenderWriteError(err error) {
	// We choose to ignore the error if we cannot log it on stderr.
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
}
