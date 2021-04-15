package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"tincho.example/security"
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

func AdminAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, error := security.GetUserFromToken(c)
		if error != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		} else {
			if token.Role == "ADMIN" {
				c.Next()
			} else {
				c.AbortWithStatusJSON(401, gin.H{"error": "the user is not admin"})
			}
		}
	}
}
