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
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[ERROR] listen: %s\n", err)
		}
	}()

	// system signal receiver
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[WARN] shutting down server...")

	// context to shutdown server
	ctxToStop, cancelCtxToStop := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	if err := srv.Shutdown(ctxToStop); err != nil {
		log.Fatal("[WARN] server forced to shutdown: ", err)
	}
	cancelCtxToStop()

	log.Println("[INFO] server exiting")
}
