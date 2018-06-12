package soba

import (
	"io/ioutil"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// A Config describes a soba instance configuration.
type Config struct {
}

// ParseConfig parses a file path and creates a new Config
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
	// TODO
	return nil
}
