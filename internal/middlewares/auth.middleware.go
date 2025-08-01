package middlewares

import (
	"learning-french-service/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrCodeUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
