package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"tincho.example/controllers"
	"tincho.example/database"
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

	var filePath string
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	} else {
		filePath = "from os vars"
	}
	e := injector.InitializeEvent(filePath)
	e.Start()

	e.GetDatabase().GetConnection().AutoMigrate(
		&database.Player{},
		&database.Category{},
	)

	addAdminUser(e)

	controllers.Config(e, r)

	r.Use(CORSMiddleware())
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://fecha-libre-view.herokuapp.com")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
