package main

import (
	"context"
	"log"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/server"
)

func main() {
	cfg := config.GetVerifiedConfig()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	dbConn, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] connecting to database : %s", err)
	}

	server.Start(
		ctx,
		cfg,
		dbConn,
	)
}
