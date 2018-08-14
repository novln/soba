package soba_test

import (
	"strings"
	"testing"

	random "github.com/Pallinder/go-randomdata"

	"github.com/novln/soba"
)

// NewHandler creates a default handler for unit test and benchmark.
func NewHandler() (soba.Handler, error) {
	return soba.Create(&soba.Config{
		Loggers: map[string]soba.ConfigLogger{},
		Appenders: map[string]soba.ConfigAppender{
			"stdout": {
				Type: "console",
			},
		},
		Root: soba.ConfigLogger{
			Level:    "disabled",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		},
	})
}

// Benchmark allocation of new Logger from Handler.
func BenchmarkHandler_NewLogger(b *testing.B) {
	handler, err := NewHandler()
	if err != nil {
		b.Fatal(err)
	}

	list := []string{}

	number := random.Number(100, 500)
	for i := 0; i < number; i++ {
		levels := []string{}
		for y := 0; y < 10; y++ {
			if (y % 2) == 0 {
				levels = append(levels, strings.ToLower(random.Noun()))
			} else {
				levels = append(levels, strings.ToLower(random.Adjective()))
			}
		}
		for y := 0; y < 10; y++ {
			if random.Number(0, 2) == 0 {
				name := strings.Join(levels[0:y], ".")
				if soba.IsLoggerNameValid(name) {
					list = append(list, name)
				}
			}
		}
	}

	max := len(list)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {

		// Create a local logger and update it's value from handler to prevent the
		// compiler to eliminate function execution.
		var l soba.Logger

		i := 0
		for pb.Next() {
			i = (i + 1) % max
			l = handler.New(list[i])
		}

		// Store logger instance in a global variable so the compiler cannot eliminate the benchmark.
		// It create a race conditions but it's okay since it's only a benchmark and not a unit test.
		gl = l

	})
}

// Test hierarchy of loggers.
func TestHandler_NewLogger(t *testing.T) {
	apiAppender := NewTestAppender("api-log")
	dbAppender := NewTestAppender("db-log")
	authAppender := NewTestAppender("auth-log")
	stdoutAppender := NewTestAppender("stdout")

	err := soba.RegisterAppenders(apiAppender, dbAppender, authAppender, stdoutAppender)
	if err != nil {
		t.Fatal(err)
	}

	handler, err := soba.Create(&soba.Config{
		Root: soba.ConfigLogger{
			Level:    "info",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		},
		Loggers: map[string]soba.ConfigLogger{
			"components.auth": {
				Level:    "info",
				Additive: true,
				Appenders: []string{
					"auth-log",
				},
			},
			"repositories": {
				Level:    "info",
				Additive: false,
				Appenders: []string{
					"db-log",
				},
			},
			"views": {
				Level:    "warning",
				Additive: false,
				Appenders: []string{
					"api-log",
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = handler

	t.Fatal("Finish this unit test")

}
