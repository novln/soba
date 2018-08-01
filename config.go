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
	Root      ConfigLogger
	Appenders map[string]ConfigAppender
	Loggers   map[string]ConfigLogger
}

// A ConfigLogger describes a logger configuration.
type ConfigLogger struct {
	Level     string
	Appenders []string
	Additive  bool
}

// A ConfigAppender describes an appender configuration.
type ConfigAppender struct {
	Type string
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
	conf.Root.Additive = false
	if !IsLevelNameValid(conf.Root.Level) {
		return errors.Errorf("level is invalid for root logger")
	}
	if len(conf.Root.Appenders) == 0 {
		return errors.Errorf("one appender is required for root logger")
	}
	for _, appender := range conf.Root.Appenders {
		_, ok := conf.Appenders[appender]
		if !ok {
			return errors.Errorf("appender is invalid for root logger")
		}
	}

	for name, logger := range conf.Loggers {
		if !IsLevelNameValid(logger.Level) {
			return errors.Errorf("level is invalid for logger: %s", name)
		}
		for _, appender := range logger.Appenders {
			_, ok := conf.Appenders[appender]
			if !ok {
				return errors.Errorf("appender is invalid for logger: %s", name)
			}
		}
	}

	for name, appender := range conf.Appenders {
		switch appender.Type {
		case "console":

		case "file":
			if appender.Path == "" {
				return errors.Errorf("path is invalid for appender: %s", name)
			}

		default:
			return errors.Errorf("type is invalid for appender: %s", name)
		}
	}

	return nil
}
