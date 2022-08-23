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

	"github.com/gin-gonic/gin"
)

type Server struct {
	Handler     *gin.Engine
	Database    database.Conn
	TokenSecret []byte
	Sesssions   *Sessions
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func Start(ctx context.Context, cfg config.Config, dbConn database.Conn) {

	srvr := Server{
		Handler:     gin.New(),
		Database:    dbConn,
		TokenSecret: []byte(cfg.TokenSecret),
		Sesssions:   NewSessions(),
	}

	// add routes
	srvr.addRoutes(cfg.Server.PathPrefix)

	// add middlewares
	srvr.addMiddlewares(
		authMiddleware(),
		middleware1(),
	)

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
			log.Fatalf("[ERROR] starting server : %s", err)
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

	if err := srv.Shutdown(ctxToStop); err != nil {
		log.Fatalf("[WARN] server forced to shutdown : %s", err)
	}
	cancelCtxToStop()

	log.Println("[INFO] server exiting")
}
