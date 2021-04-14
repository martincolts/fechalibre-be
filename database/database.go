package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	var database = &Database{}
	database.ConfigDb()
	return database
}

func (database *Database) ConfigDb() {
	dsn := "host=localhost user=testing password=testing dbname=testing port=5432"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database.db = db
}

func (database *Database) GetConnection() *gorm.DB {
	return database.db
}
