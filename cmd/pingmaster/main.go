package main

import (
	"context"

	"pingmaster/config"
	"pingmaster/server"
)

func main() {
	cfg := config.GetVerifiedConfig()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	server.Start(ctx, cfg.Server)
}
