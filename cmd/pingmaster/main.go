package main

import (
	"context"
	"log"
	"time"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/server"
	"pingmaster/target/targetspool"
)

func main() {
	cfg := config.GetVerifiedConfig()

	ctx, cancelCtx := context.WithCancel(context.Background())

	dbConn, err := database.New(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("[ERR] connecting to database : %s", err)
	}

	// fetch targets from DB
	targets, err := dbConn.FetchTargets(ctx)
	if err != nil {
		log.Fatalf("[ERR] fetching targets from database : %s", err)
	}

	targetsPool := targetspool.New()

	// add targets from DB to pool
	addCount := 0
	for _, eachTarget := range targets {
		err = targetsPool.Add(eachTarget)
		if err != nil {
			log.Printf(
				"[ERR] adding target from DB to pool : %s, target : %+v",
				err, eachTarget,
			)
			continue
		}
		addCount++
	}
	log.Printf("[INF] added %d targets from DB to pool", addCount)

	// Monitor targets pool in separate goroutine
	go targetsPool.Monitor(ctx, dbConn)

	server.Start(
		ctx,
		cfg,
		dbConn,
		targetsPool,
	)
	cancelCtx()

	log.Println("[INF] stopping pingmaster")

	// wait for all resources to get released
	time.Sleep(time.Second * 2)
}
