package controllers

import (
	"fmt"

	"tincho.example/controllers/middlewares"
	"tincho.example/injector"
	"tincho.example/security"
	"tincho.example/services"

	"github.com/gin-gonic/gin"
)

func Config(e *injector.Event, r *gin.Engine) {
	fmt.Println("loading config")

	public := r.Group("/public")
	{
		public.GET("/ping", services.Ping())
		public.POST("/login", login(e))
	}

	private := r.Group("/internal")
	private.Use(middlewares.Authorize())
	{
		private.GET("/validToken", services.Ping())
		private.POST("/category", InsertCategory(e))
		private.POST("/player", insertPlayer(e))
		private.GET("/player/:id", getPlayerById(e))
		private.GET("/player", getAllPlayers(e))
	}

	admin := r.Group("/admin")
	admin.Use(middlewares.AdminAuthorize())
	{
		admin.GET("/verify", func(c *gin.Context) {
			user, _ := security.GetUserFromToken(c)
			c.JSON(200, gin.H{"data": "you are admin", "user": user})
		})
	}
}
