package main

import (
	"context"
	"log"
	"time"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/server"
	"pingmaster/target"
)

func main() {
	cfg := config.GetVerifiedConfig()

	ctx, cancelCtx := context.WithCancel(context.Background())

	dbConn, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("[ERROR] connecting to database : %s", err)
	}

	targetsPool := target.NewPool()
	// Monitor targets pool in separate goroutine
	go targetsPool.Monitor(ctx)

	server.Start(
		ctx,
		cfg,
		dbConn,
		targetsPool,
	)
	cancelCtx()

	log.Println("[INFO] stopping pingmaster")

	// wait for all resources to get released
	time.Sleep(time.Second * 2)
}
