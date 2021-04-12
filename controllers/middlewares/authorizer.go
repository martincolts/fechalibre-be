package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header.Get("Authorization")

		token := strings.Split(headers, " ")[1]

		if token == "valid" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "not authorized",
			})
		}

	}
}
