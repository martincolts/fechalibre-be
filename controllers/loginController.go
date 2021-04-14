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
		player, _ := playerService.GetPlayerByUsername(loginForm.Username)
		if player.Password == loginForm.Password {
			tokenString, _ := security.CreateToken(player)
			c.JSON(200, gin.H{"token": tokenString})
		}
	}
}
