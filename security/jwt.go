package security

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"tincho.example/database"
)

var mySecret string = "fechaLibreSecret"

type MyCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func CreateToken(player *database.Player) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["username"] = player.Username
	atClaims["role"] = player.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	signedToken, _ := token.SignedString([]byte(mySecret))

	return signedToken, nil
}

func GetUserFromToken(c *gin.Context) (*MyCustomClaims, error) {
	headers := c.Request.Header.Get("Authorization")
	tokenString := strings.Split(headers, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &MyCustomClaims{
			Username: claims["username"].(string),
			Role:     claims["role"].(string),
		}, nil
	}

	return nil, errors.New("invalid token")
}
