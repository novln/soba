package main

import (
	"context"

	"github.com/novln/soba"
)

func main() {
	ctx := context.Background()

	ctx, err := soba.LoadWithFile(ctx, "conf.yaml")
	if err != nil {
		panic(err)
	}

	ctx, err = soba.Load(ctx)
	if err != nil {
		panic(err)
	}

	logger := soba.New(ctx, "app.store.users")
	logger.Info("Hello", soba.String("ok", "fuck"))

}
