package soba

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// DefaultConfigPath defines the default configuration path.
const DefaultConfigPath = "soba.yml"

// A Config describes a soba instance configuration.
type Config struct {
	// Root defines default logger configuration.
	Root ConfigLogger
	// Appenders is a list of appenders definition.
	Appenders map[string]ConfigAppender
	// Loggers is a list of loggers configuration.
	Loggers map[string]ConfigLogger
	// verified defines if configuration has been validated.
	verified bool
}

// IsAppenderExists verifies if appender identified by given name exists.
func (conf *Config) IsAppenderExists(name string) bool {
	_, ok := conf.Appenders[name]
	if ok {
		return true
	}

	plMutex.Lock()
	defer plMutex.Unlock()

	_, ok = plAppenders[name]
	return ok
}

// A ConfigLogger describes a logger configuration.
type ConfigLogger struct {
	Level     string
	Appenders []string
	Additive  bool
}

// A ConfigAppender describes an appender configuration.
type ConfigAppender struct {
	// Type defines an appender type. Could be "console" or "file".
	Type string
	// Path defines the file path for a file appender.
	Path string
}

// CheckPath verifies that given path is valid.
func CheckPath(path string) bool {
	file, err := os.Stat(path)
	return err == nil && !file.IsDir()
}

// ParseConfig parses a file path and creates a new Config.
func ParseConfig(path string) (*Config, error) {
	conf := &Config{}

	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read configuration file")
	}

	err = yaml.Unmarshal(buffer, conf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse configuration file")
	}

	err = ValidateConfig(conf)
	if err != nil {
		return nil, errors.Wrap(err, "configuration file is invalid")
	}

	return conf, nil
}

// ValidateConfig verifies that given configuration is valid.
func ValidateConfig(conf *Config) error {
	if conf == nil {
		return errors.Errorf("given configuration is empty")
	}
	if conf.verified {
		return nil
	}

	err := validateAppendersConfig(conf)
	if err != nil {
		return err
	}

	err = validateRootLoggerConfig(conf)
	if err != nil {
		return err
	}

	err = validateLoggersConfig(conf)
	if err != nil {
		return err
	}

	conf.verified = true

	return nil
}

func validateRootLoggerConfig(conf *Config) error {

	conf.Root.Additive = false

	if !IsLevelNameValid(conf.Root.Level) {
		return errors.Errorf("level is invalid for root logger")
	}

	if len(conf.Root.Appenders) == 0 {
		return errors.Errorf("one appender is required for root logger")
	}

	for _, appender := range conf.Root.Appenders {
		if !conf.IsAppenderExists(appender) {
			return errors.Errorf("appender is invalid for root logger")
		}
	}

	return nil
}

func validateLoggersConfig(conf *Config) error {

	for name, logger := range conf.Loggers {

		if !IsLevelNameValid(logger.Level) {
			return errors.Errorf("level is invalid for logger: %s", name)
		}

		if !IsLoggerNameValid(name) {
			return errors.Errorf("name is invalid for logger: %s", name)
		}

		for _, appender := range logger.Appenders {
			if !conf.IsAppenderExists(appender) {
				return errors.Errorf("appender is invalid for logger: %s", name)
			}
		}

	}

	return nil
}

func validateAppendersConfig(conf *Config) error {

	for name, appender := range conf.Appenders {

		if !IsAppenderNameValid(name) {
			return errors.Errorf("name is invalid for appender: %s", name)
		}

		switch appender.Type {
		case ConsoleAppenderType:

		case FileAppenderType:
			if appender.Path == "" {
				return errors.Errorf("path is invalid for appender: %s", name)
			}

		default:
			return errors.Errorf("type is invalid for appender: %s", name)
		}
	}

	return nil
}
