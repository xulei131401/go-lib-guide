package middleware

import (
	"github.com/gin-gonic/gin"
)

func ParameterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*c.JSON(500, gin.H{
			"code": 500,
			"msg": "失败",
			"data": []string{},
		})
		c.Abort()*/
		c.Next()
		return
	}
}