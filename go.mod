module tincho.example

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/google/wire v0.5.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.6
)

// +heroku goVersion go1.16
// +build wireinject
