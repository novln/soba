package soba

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// A Handler provides an alternative way to obtain loggers if the context based approach doesn't
// fit your requirements.
type Handler interface {
	// New creates a new Logger using given name.
	New(name string) Logger
}

// Create provides an alternative way to obtain loggers if the context based approach doesn't
// fit your requirements by creating a handler instance from given configuration.
func Create(conf *Config) (Handler, error) {
	err := ValidateConfig(conf)
	if err != nil {
		return nil, errors.Wrap(err, "configuration is invalid")
	}
	return create(conf)
}

// A handler contains every required components to provides loggers.
type handler struct {
	conf      Config
	appenders map[string]Appender
	loggers   map[string]*logger
}

func create(conf *Config) (*handler, error) {
	handler := &handler{
		conf:      *conf,
		appenders: map[string]Appender{},
		loggers:   map[string]*logger{},
	}

	err := createAppenders(conf, handler)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create soba handler")
	}

	err = createRootLogger(conf, handler)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create soba handler")
	}

	fmt.Printf("%+v\n", conf.Root)

	return handler, nil
}

func createAppenders(conf *Config, handler *handler) error {
	for name := range conf.Appenders {
		appender, err := NewAppender(conf.Appenders[name])
		if err != nil {
			return err
		}
		handler.appenders[name] = appender
	}
	return nil
}

func createRootLogger(conf *Config, handler *handler) error {
	appenders := []Appender{}

	for _, name := range conf.Root.Appenders {
		appender, ok := handler.appenders[name]
		if !ok {
			return errors.Errorf("unknown appender name: %s", name)
		}
		appenders = append(appenders, appender)
	}

	level, ok := ParseLevel(conf.Root.Level)
	if !ok {
		return errors.Errorf("unknown level name: %s", conf.Root.Level)
	}

	logger := &logger{
		level:     level,
		appenders: appenders,
	}

	handler.loggers[""] = logger

	return nil
}

func (h *handler) New(name string) Logger {
	// Try to find a logger identified by given name.
	logger, ok := h.loggers[name]
	if ok {
		return logger
	}

	// Otherwise, try to find a ancestor one.
	hierarchy := strings.Split(name, ".")
	length := len(hierarchy)
	for i := 1; i < length; i++ {
		cursor := length - i
		list := hierarchy[0:cursor]
		current := strings.Join(list, ".")

		logger, ok := h.loggers[current]
		if ok {
			return logger.withName(name)
		}
	}

	// Finally, returns root logger.
	return h.loggers[""].withName(name)
}
