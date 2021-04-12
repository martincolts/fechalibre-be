package controllers

import (
	"fmt"

	"tincho.example/controllers/middlewares"
	"tincho.example/services"

	"github.com/gin-gonic/gin"
)

func Config(r *gin.Engine) {
	fmt.Println("loading config")

	public := r.Group("/public")
	{
		public.GET("/ping", services.Ping())
	}

	private := r.Group("/internal")
	private.Use(middlewares.Authorize())
	{
		private.GET("/validToken", services.Ping())
	}
}
