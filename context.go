package soba

import (
	"context"
	"os"
)

type ctxKey struct{}

// Load returns a new context with a soba instance. It relies on conventions and default configurations:
//  - First, it will lookup from environment variable if a configuration path is defined.
//  - Then, it will from current directory if a configuration file exists.
//  - Finally, it will create a new instance with default configurations.
//
// For specific configurations, please uses either LoadWithConfig or LoadWithFile.
func Load(ctx context.Context) (context.Context, error) {
	path := os.Getenv("SOBA_CONF")
	if path != "" && CheckPath(path) {
		return LoadWithFile(ctx, path)
	}

	if CheckPath(DefaultConfigPath) {
		return LoadWithFile(ctx, DefaultConfigPath)
	}

	return LoadWithConfig(ctx, Config{
		Loggers: map[string]ConfigLogger{},
		Appenders: map[string]ConfigAppender{
			"stdout": ConfigAppender{
				Type: "console",
			},
		},
		Root: ConfigLogger{
			Level:    "info",
			Additive: false,
			Appenders: []string{
				"stdout",
			},
		},
	})
}

// LoadWithConfig returns a new context with a soba instance using given configuration.
func LoadWithConfig(ctx context.Context, config Config) (context.Context, error) {
	// TODO
	return ctx, nil
}

// LoadWithFile returns a new context with a soba instance using given file path.
func LoadWithFile(ctx context.Context, path string) (context.Context, error) {
	conf, err := ParseConfig(path)
	if err != nil {
		return ctx, err
	}
	return LoadWithConfig(ctx, *conf)
}
