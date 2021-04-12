package controllers

import (
	"github.com/gin-gonic/gin"
	"tincho.example/database"
	"tincho.example/injector"
)

func InsertCategory(e *injector.Event) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category database.Category
		c.BindJSON(&category)
		error, result := e.GetCategoryService().Insert(&category)
		if error != nil {
			c.JSON(409, gin.H{"error": error})
		} else {
			c.JSON(200, gin.H{"data": result})
		}
	}
}
