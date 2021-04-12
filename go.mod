module tincho.example

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/wire v0.5.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.6
)

// +heroku goVersion go1.16
// +build wireinject