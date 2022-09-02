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
	"pingmaster/target"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Handler     *gin.Engine
	Database    database.Conn
	TokenSecret []byte
	Sesssions   *Sessions
	TargetsPool *target.Pool
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Start(ctx context.Context, cfg config.Config, dbConn database.Conn, targetsPool *target.Pool) {

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
			"[INFO] starting server on port : %d",
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
			log.Printf("[ERROR] starting server : %s", err)
			return
		}
	}()

	// system signal receiver
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[WARN] shutting down server")

	// context to shutdown server
	ctxToStop, cancelCtxToStop := context.WithTimeout(
		ctx,
		time.Second*5,
	)
	defer cancelCtxToStop()

	if err := srv.Shutdown(ctxToStop); err != nil {
		log.Printf("[WARN] server forced to shutdown : %s", err)
		return
	}

	log.Println("[INFO] server exiting")
}
