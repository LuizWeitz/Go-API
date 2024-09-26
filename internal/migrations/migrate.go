package migrations

import (
	"log"

	"github.com/luizweitz/go-api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrateAll(db *gorm.DB) {
	if err := db.Table("users").AutoMigrate(&models.User{}); err != nil {
		log.Printf("Warning: could not migrate table users: %v", err)
	}
}
