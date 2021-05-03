package middlewares

import (
	"github.com/gin-gonic/gin"
	"tincho.example/security"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, error := security.GetUserFromToken(c)
		if error != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		} else {
			if token.Role == "ADMIN" || token.Role == "PLAYER" {
				c.Next()
			} else {
				c.AbortWithStatusJSON(401, gin.H{"error": "the user is not admin"})
			}
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
