package main

import (
	"fmt"
	"net/http"
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

	r.Use(CORS())
	r.Use(CORSMiddleware())

	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Referrer Policy")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func CORS() func(c *gin.Context) {
	return func(c *gin.Context) {

		// First, we add the headers with need to enable CORS
		// Make sure to adjust these headers to your needs
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Content-Type", "application/json")

		// Second, we handle the OPTIONS problem
		if c.Request.Method != "OPTIONS" {

			c.Next()

		} else {

			// Everytime we receive an OPTIONS request,
			// we just return an HTTP 200 Status Code
			// Like this, Angular can now do the real
			// request using any other method than OPTIONS
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
