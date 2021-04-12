package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigDb() (db *gorm.DB, err error) {
	dsn := "host=localhost user=testing password=testing dbname=testing port=5432"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
