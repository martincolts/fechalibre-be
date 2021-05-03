package controllers

import (
	"github.com/gin-gonic/gin"
	"tincho.example/security"
)

func getMenuItems() func(c *gin.Context) {
	return func(c *gin.Context) {
		if permissions, error := permissions(c); error == nil {
			c.JSON(200, gin.H{"result": permissions})
		} else {
			c.JSON(401, gin.H{"error": "unauthorized"})
		}
	}
}

func permissions(c *gin.Context) ([]string, error) {
	permissions := make(map[string][]string)
	permissions["ADMIN"] = []string{"Crear Jugador", "Editar mi Perfil"}
	permissions["PLAYER"] = []string{"Editar mi Perfil"}

	if token, error := security.GetUserFromToken(c); error == nil {
		return permissions[token.Role], nil
	} else {
		return nil, error
	}
}
