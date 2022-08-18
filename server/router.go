package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(handler *gin.Engine) {

	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
