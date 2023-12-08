package middleware

import (
	"github.com/gin-gonic/gin"
)

func DevMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}