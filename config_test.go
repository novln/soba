package soba_test

import (
	"testing"

	"github.com/novln/soba"
)

// Test path verification for a configuration file.
func TestConfig_CheckPath(t *testing.T) {
	{
		path := "testdata/foobar.yaml"
		value := soba.CheckPath(path)
		if value {
			t.Fatalf(`Unexpected result: "%s" should not be a valid path`, path)
		}
	}
	{
		path := "testdata/simple.yaml"
		value := soba.CheckPath(path)
		if !value {
			t.Fatalf(`Unexpected result: "%s" should be a valid path`, path)
		}
	}
	{
		path := "testdata/invalid.yaml"
		value := soba.CheckPath(path)
		if !value {
			t.Fatalf(`Unexpected result: "%s" should be a valid path`, path)
		}
	}
	{
		path := "examples"
		value := soba.CheckPath(path)
		if value {
			t.Fatalf(`Unexpected result: "%s" should not be a valid path`, path)
		}
	}
}

// Test parsing of a configuration file.
// nolint: gocyclo
func TestConfig_ParseConfig(t *testing.T) {
	{
		path := "testdata/simple.yaml"
		conf, err := soba.ParseConfig(path)
		if err != nil {
			t.Fatalf("Unexpected error for %s: %+v", path, err)
		}
		if conf == nil {
			t.Fatalf("A configuration was expected for %s", path)
		}
	}
	{
		path := "testdata/invalid.yaml"
		conf, err := soba.ParseConfig(path)
		if err == nil {
			t.Fatalf("An error was expected for %s", path)
		}
		if conf != nil {
			t.Fatalf("Unexpected configuration for %s", path)
		}
	}
	{
		path := "testdata/foobar.yaml"
		conf, err := soba.ParseConfig(path)
		if err == nil {
			t.Fatalf("An error was expected for %s", path)
		}
		if conf != nil {
			t.Fatalf("Unexpected configuration for %s", path)
		}
	}
	{
		path := "config.go"
		conf, err := soba.ParseConfig(path)
		if err == nil {
			t.Fatalf("An error was expected for %s", path)
		}
		if conf != nil {
			t.Fatalf("Unexpected configuration for %s", path)
		}
	}
	{
		path := "examples"
		conf, err := soba.ParseConfig(path)
		if err == nil {
			t.Fatalf("An error was expected for %s", path)
		}
		if conf != nil {
			t.Fatalf("Unexpected configuration for %s", path)
		}
	}
}

// Test validation of a configuration file.
// nolint: gocyclo
func TestConfig_ValidateConfig(t *testing.T) {
	{
		conf := soba.NewDefaultConfig()
		err := soba.ValidateConfig(conf)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		err = soba.ValidateConfig(conf)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
	}
	{
		err := soba.ValidateConfig(nil)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Appenders["Foobar"] = soba.ConfigAppender{
			Type: soba.ConsoleAppenderType,
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Appenders["demo"] = soba.ConfigAppender{
			Type: soba.FileAppenderType,
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Appenders["demo"] = soba.ConfigAppender{
			Type: "foobar",
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Root.Level = "foobar"

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Root.Appenders = []string{}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Root.Appenders = []string{
			"foobar",
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Loggers["Foobar"] = soba.ConfigLogger{
			Level:    soba.InfoLevel.String(),
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Loggers["demo"] = soba.ConfigLogger{
			Level:    "foobar",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
	{
		conf := soba.NewDefaultConfig()
		conf.Loggers["demo"] = soba.ConfigLogger{
			Level:    soba.InfoLevel.String(),
			Additive: false,
			Appenders: []string{
				"foobar",
			},
		}

		err := soba.ValidateConfig(conf)
		if err == nil {
			t.Fatal("An error was expected")
		}
	}
}
