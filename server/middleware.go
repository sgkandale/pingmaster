package server

import "github.com/gin-gonic/gin"

func (s Server) addMiddlewares(middlewares ...gin.HandlerFunc) {
	s.Handler.Use(
		middlewares...,
	)
}

func middleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
