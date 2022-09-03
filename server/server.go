package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pingmaster/config"
	"pingmaster/database"
	"pingmaster/target/targetspool"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Handler     *gin.Engine
	Database    database.Conn
	TokenSecret []byte
	Sesssions   *Sessions
	TargetsPool *targetspool.Pool
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Start(ctx context.Context, cfg config.Config, dbConn database.Conn, targetsPool *targetspool.Pool) {

	srvr := &Server{
		Handler:     gin.New(),
		TokenSecret: []byte(cfg.Security.TokenSecret),
		Database:    dbConn,
		Sesssions:   NewSessions(),
		TargetsPool: targetsPool,
	}

	// add middlewares
	srvr.Handler.Use(
		headers(cfg.Security),
		authMiddleware(),
	)

	// add routes
	srvr.addRoutes(cfg.Server.PathPrefix)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: srvr.Handler,
	}

	go func() {
		log.Printf(
			"[INF] starting server on port : %d",
			cfg.Server.Port,
		)

		var err error
		if cfg.Server.TLS {
			err = srv.ListenAndServeTLS(
				cfg.Server.CertPath,
				cfg.Server.KeyPath,
			)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Printf("[ERR] starting server : %s", err)
			return
		}
	}()

	// system signal receiver
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[WRN] shutting down server")

	// context to shutdown server
	ctxToStop, cancelCtxToStop := context.WithTimeout(
		ctx,
		time.Second*5,
	)
	defer cancelCtxToStop()

	if err := srv.Shutdown(ctxToStop); err != nil {
		log.Printf("[WRN] server forced to shutdown : %s", err)
		return
	}

	log.Println("[INF] server exiting")
}
