package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {

	var err error

	dsn := os.Getenv("DATABASE_URL")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect to DataBase")
	}

	return db
}
