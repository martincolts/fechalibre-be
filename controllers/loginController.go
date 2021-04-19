package controllers

import (
	"github.com/gin-gonic/gin"
	"tincho.example/injector"
	"tincho.example/security"
)

type Login struct {
	Username string
	Password string
}

func login(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		playerService := e.GetPlayerService()
		var loginForm Login
		c.BindJSON(&loginForm)
		if player, error := playerService.GetPlayerByUsername(loginForm.Username); error != nil {
			c.JSON(401, gin.H{"error": "El usuario no existe"})
		} else if player.Password == loginForm.Password {
			tokenString, _ := security.CreateToken(player)
			c.JSON(200, gin.H{"token": tokenString})
		} else {
			c.JSON(401, gin.H{"error": "Pusiste la clave mal gil"})
		}
	}
}
