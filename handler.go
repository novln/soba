package soba

import (
	"strings"
	"sync"

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
	loggers   sync.Map
}

func create(conf *Config) (*handler, error) {
	handler := &handler{
		conf:      *conf,
		appenders: map[string]Appender{},
		loggers:   sync.Map{},
	}

	err := createAppenders(conf, handler)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create soba handler")
	}

	err = createRootLogger(conf, handler)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create soba handler")
	}

	err = createChildLoggers(conf, handler)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create soba handler")
	}

	return handler, nil
}

func createAppenders(conf *Config, handler *handler) error {
	for name := range conf.Appenders {
		appender, err := NewAppender(name, conf.Appenders[name])
		if err != nil {
			return err
		}
		handler.appenders[name] = appender
	}
	return nil
}

func createRootLogger(conf *Config, handler *handler) error {
	level, err := getLoggerLevel(conf.Root, "root")
	if err != nil {
		return err
	}

	appenders, err := getAppendersForRootLogger(conf, handler)
	if err != nil {
		return err
	}

	handler.loggers.Store("", &logger{
		level:     level,
		appenders: appenders,
	})

	return nil
}

func getLoggerLevel(conf ConfigLogger, name string) (Level, error) {
	level, ok := ParseLevel(conf.Level)
	if !ok {
		return UnknownLevel, errors.Errorf("unknown level for logger '%s': %s", name, conf.Level)
	}
	return level, nil
}

func getParentAppendersForLogger(conf *Config, handler *handler, hierarchy []string, result map[string]Appender) error {

	length := len(hierarchy)
	for i := 1; i < length; i++ {

		cursor := length - i
		list := hierarchy[0:cursor]
		current := strings.Join(list, ".")

		parent, ok := conf.Loggers[current]
		if ok {
			return getLocalAppendersForLogger(parent, handler, result)
		}
	}

	return nil
}

func getLocalAppendersForLogger(conf ConfigLogger, handler *handler, result map[string]Appender) error {
	for _, name := range conf.Appenders {
		appender, ok := handler.appenders[name]
		if !ok {
			return errors.Errorf("unknown appender name: '%s'", name)
		}
		result[name] = appender
	}
	return nil
}

func getAppendersForRootLogger(conf *Config, handler *handler) ([]Appender, error) {
	appenders := map[string]Appender{}

	getLocalAppendersForLogger(conf.Root, handler, appenders)
	list := []Appender{}
	for _, appender := range appenders {
		list = append(list, appender)
	}

	return list, nil
}

func getAppendersForChildLogger(conf *Config, handler *handler, name string) ([]Appender, error) {
	appenders := map[string]Appender{}

	getLocalAppendersForLogger(conf.Loggers[name], handler, appenders)
	if conf.Loggers[name].Additive {
		hierarchy := strings.Split(name, ".")
		getParentAppendersForLogger(conf, handler, hierarchy, appenders)
	}

	list := []Appender{}
	for _, appender := range appenders {
		list = append(list, appender)
	}

	return list, nil
}

func createChildLoggers(conf *Config, handler *handler) error {
	for name := range conf.Loggers {
		level, err := getLoggerLevel(conf.Loggers[name], name)
		if err != nil {
			return err
		}

		appenders, err := getAppendersForChildLogger(conf, handler, name)
		if err != nil {
			return err
		}

		handler.loggers.Store(name, &logger{
			level:     level,
			appenders: appenders,
		})

	}
	return nil
}

func (h *handler) New(name string) Logger {
	// First, try to find the logger identified by given name.
	val, ok := h.loggers.Load(name)
	if ok {
		return val.(*logger)
	}

	// Next, try to find a ancestor one by moving up to the hierarchy.
	hierarchy := strings.Split(name, ".")
	length := len(hierarchy)
	for i := 1; i < length; i++ {
		cursor := length - i
		list := hierarchy[0:cursor]
		current := strings.Join(list, ".")

		val, ok = h.loggers.Load(current)
		if ok {
			copy := val.(*logger).copyWithName(name)
			val, _ = h.loggers.LoadOrStore(name, copy)
			return val.(*logger)
		}
	}

	// Finally, use root logger as default.
	val, ok = h.loggers.Load("")
	if !ok {
		panic("soba: root logger must be defined")
	}

	copy := val.(*logger).copyWithName(name)
	val, _ = h.loggers.LoadOrStore(name, copy)
	return val.(*logger)
}
