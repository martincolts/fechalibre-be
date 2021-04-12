package main

import (
	"fmt"
	"time"

	"tincho.example/controllers"
	"tincho.example/database"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hola mundo")
	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controllers.Config(r)
	db, err := database.ConfigDb()
	if err != nil {
		fmt.Errorf("this was an error to connect to the database", err)
	}
	db.AutoMigrate(&database.Player{})
	db.AutoMigrate(&database.Category{})
	// var category = database.Category{Name: "Another name"}
	// result := db.Create(&category)

	var player = &database.Player{
		Created:   time.Now().Unix(),
		Birthdate: 0,
		Name:      "Martin",
		Lastname:  "Lopez",
		DNI:       "34421478",
	}

	result := db.Create(player)

	fmt.Println("ERROR", result.Error)
	r.Run()

}
