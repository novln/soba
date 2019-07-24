package soba_test

import (
	"context"
	"os"
	"testing"

	"github.com/novln/soba"
)

// Test creation of a new handler with context.
// nolint: gocyclo
func TestContext_Load(t *testing.T) {

	checkSuccess := func(ctx context.Context) {
		defer func() {
			thr := recover()
			if thr != nil {
				t.Fatalf("Unexpected panic: %+v", thr)
			}
		}()
		// Ensure we don't have panic while creating a new logger.
		soba.New(ctx, "foobar")
	}

	checkFailure := func(ctx context.Context) {
		defer func() {
			thr := recover()
			if thr == nil {
				t.Fatal("A panic error was expected")
			}
		}()
		// Ensure we don't have panic while creating a new logger.
		soba.New(ctx, "foobar")
	}

	{
		// Create a handler using "SOBA_CONF" environment variable with a valid file.
		ctx := context.Background()
		env := os.Getenv(soba.EnvConfigPath)

		err := os.Setenv(soba.EnvConfigPath, "testdata/simple.yaml")
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		ctx, err = soba.Load(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		err = os.Setenv(soba.EnvConfigPath, env)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		defer checkSuccess(ctx)
	}
	{
		// Create a handler using "SOBA_CONF" environment variable with an invalid file.
		ctx := context.Background()
		env := os.Getenv(soba.EnvConfigPath)

		err := os.Setenv(soba.EnvConfigPath, "config.go")
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		ctx, err = soba.Load(ctx)
		if err == nil {
			t.Fatalf("An error was expected")
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		err = os.Setenv(soba.EnvConfigPath, env)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		defer checkFailure(ctx)
	}
	{
		// Create a handler without "SOBA_CONF" environment variable and
		// without a default configuration file.
		ctx := context.Background()
		env := os.Getenv(soba.EnvConfigPath)

		err := os.Setenv(soba.EnvConfigPath, "")
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		ctx, err = soba.Load(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		err = os.Setenv(soba.EnvConfigPath, env)
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}

		defer checkSuccess(ctx)
	}
	{
		// Create a handler with an invalid configuration file path.
		ctx := context.Background()
		ctx, err := soba.LoadWithFile(ctx, "foobar.yml")
		if err == nil {
			t.Fatalf("An error was expected")
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		defer checkFailure(ctx)
	}
	{
		// Create a handler with an invalid configuration file format.
		ctx := context.Background()
		ctx, err := soba.LoadWithFile(ctx, "config.go")
		if err == nil {
			t.Fatalf("An error was expected")
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		defer checkFailure(ctx)
	}
	{
		// Create a handler with a valid configuration file.
		ctx := context.Background()
		ctx, err := soba.LoadWithFile(ctx, "testdata/simple.yaml")
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		defer checkSuccess(ctx)
	}
	{
		// Create a handler with a wrong configuration.
		ctx := context.Background()
		ctx, err := soba.LoadWithConfig(ctx, &soba.Config{
			Root: soba.ConfigLogger{
				Level:    "info",
				Additive: false,
				Appenders: []string{
					"stdout",
				},
			},
		})
		if err == nil {
			t.Fatalf("An error was expected")
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		defer checkFailure(ctx)
	}
	{
		// Create a handler with a correct configuration.
		ctx := context.Background()
		ctx, err := soba.LoadWithConfig(ctx, &soba.Config{
			Root: soba.ConfigLogger{
				Level:    "info",
				Additive: false,
				Appenders: []string{
					"stdout",
				},
			},
			Loggers: map[string]soba.ConfigLogger{},
			Appenders: map[string]soba.ConfigAppender{
				"stdout": {
					Type: "console",
				},
			},
		})
		if err != nil {
			t.Fatalf("Unexpected error: %+v", err)
		}
		if ctx == nil {
			t.Fatal("A context was expected")
		}

		defer checkSuccess(ctx)
	}
}
