package soba

import (
	"context"
)

type ctxKey struct{}

// Load returns a new context with a soba instance. It relies on conventions and default configurations.
// For specific configurations, please uses either LoadWithConfig or LoadWithFile.
func Load(ctx context.Context) (context.Context, error) {
	// TODO
	return ctx, nil
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
