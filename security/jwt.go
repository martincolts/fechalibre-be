package security

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"tincho.example/database"
)

var mySecret string = "fechaLibreSecret"

type MyCustomClaims struct {
	jwt.StandardClaims
	Username string              `json:"username"`
	Role     database.PlayerRole `json:"role"`
}

func CreateToken(player *database.Player) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["username"] = player.Username
	atClaims["role"] = player.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	signedToken, _ := token.SignedString([]byte(mySecret))

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
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
	return token, nil
}
