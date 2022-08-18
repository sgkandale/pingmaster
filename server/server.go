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
	Handler  *gin.Engine
	Database database.Database
}

func Start(ctx context.Context, cfg config.ServerConfig) {

	// gin handler
	handler := gin.New()
	gin.SetMode(gin.ReleaseMode)

	addRoutes(handler)

	// add middlewares
	handler.Use(
		middleware1(),
	)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}

	go func() {
		log.Printf(
			"[INFO] starting server on port : %d",
			cfg.Port,
		)

		var err error
		if cfg.TLS {
			err = srv.ListenAndServeTLS(
				cfg.CertPath,
				cfg.KeyPath,
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
		5*time.Second,
	)

	if err := srv.Shutdown(ctxToStop); err != nil {
		log.Fatalf("[WARN] server forced to shutdown : %s", err)
	}
	cancelCtxToStop()

	log.Println("[INFO] server exiting")
}
