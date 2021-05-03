package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tincho.example/database"
	"tincho.example/injector"
)

func insertPlayer(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		var player database.Player
		ok := c.BindJSON(&player)
		if ok != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "bad request"})
		} else {
			result, error := e.GetPlayerService().Insert(&player)
			if error != nil {
				c.JSON(409, gin.H{"error": error})
				return
			} else {
				c.JSON(200, gin.H{"data": result})
				return
			}
		}
	}
}

func getAllPlayers(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		if players, error := e.GetPlayerService().GetAllPlayers(); error == nil {
			c.JSON(200, gin.H{"data": players})
		} else {
			c.JSON(409, gin.H{"errors": error})
		}

	}
}

func getPlayerById(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if intId, error := strconv.Atoi(id); error != nil {
			c.JSON(400, gin.H{"errors": "the Id is not an string"})
		} else {
			if players, error := e.GetPlayerService().GetPlayerById(intId); error == nil {
				c.JSON(200, gin.H{"data": players})
			} else {
				c.JSON(409, gin.H{"errors": error})
			}
		}
	}
}
