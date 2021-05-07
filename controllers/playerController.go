package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"tincho.example/database"
	"tincho.example/dtos"
	"tincho.example/injector"
	"tincho.example/security"
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

func updatePassword(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		playerService := e.GetPlayerService()

		var payload dtos.UpdatePasswordPayload
		ok := c.BindJSON(&payload)
		if ok != nil {
			c.JSON(400, gin.H{"errors": "payload is not correct"})
		}

		token, error := security.GetUserFromToken(c)
		if error != nil {
			c.AbortWithStatusJSON(400, gin.H{"errors": "token is not correct"})
		}

		if player, error := playerService.UpdatePassword(payload, token.Username); error != nil {
			c.AbortWithStatusJSON(500, gin.H{"errors": "error updating player"})
		} else {
			c.JSON(200, gin.H{"data": player})
			return
		}

	}
}

func getPlayerById(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if intId, error := strconv.ParseInt(id, 10, 64); error != nil {
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
