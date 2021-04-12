package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"tincho.example/controllers"
	"tincho.example/injector"
)

func main() {
	fmt.Println("hola mundo")
	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	e := injector.InitializeEvent()
	e.Start()

	controllers.Config(e, r)
	r.Run()

}
