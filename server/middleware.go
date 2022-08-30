package server

import (
	"net/http"

	"pingmaster/config"

	"github.com/gin-gonic/gin"
)

func headers(cfg config.SecurityConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, HEAD")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(
				http.StatusOK,
			)
		} else {
			c.Next()
		}
	}
}

// authMiddleware only checks for existance of the header,
// verification is handled in each handler
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authTokenArr := c.Request.Header["Authorization"]
		if authTokenArr == nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				ServerResponse{
					Status:  ResponseStatus_Unauthorized,
					Message: ResponseMessage_NoAuthHeader,
				},
			)
			return
		}
	}
}
