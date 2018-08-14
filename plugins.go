package soba

import (
	"sync"

	"github.com/pkg/errors"
)

// These global variables are only used to provide extensibility with external appenders.
// If you think of a better solution, please submit a pull request, because I hate global variables.
var (
	// plMutex is a mutex to manage concurrent access when registering an appender.
	plMutex = sync.Mutex{}
	// plAppenders is a list of external appenders identified by their name.
	plAppenders = map[string]Appender{}
)

// RegisterAppenders registers given external appenders to be accessible for loggers.
func RegisterAppenders(appenders ...Appender) error {
	plMutex.Lock()
	defer plMutex.Unlock()

	for _, appender := range appenders {

		if !IsAppenderNameValid(appender.Name()) {
			return errors.Errorf("name is invalid for appender: %s", appender.Name())
		}

		plAppenders[appender.Name()] = appender

	}

	return nil
}
