package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tincho.example/conf"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(conf *conf.Conf) *Database {
	var database = &Database{}
	database.ConfigDb(conf)
	return database
}

func (database *Database) ConfigDb(conf *conf.Conf) {
	databaseConfig := conf.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", databaseConfig.Host, databaseConfig.User,
		databaseConfig.Password, databaseConfig.Dbname, databaseConfig.Port)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database.db = db
}

func (database *Database) GetConnection() *gorm.DB {
	return database.db
}
