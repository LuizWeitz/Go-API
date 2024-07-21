package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"ID"`
	Name      string         `binding:"required" json:"name"`
	Email     string         `binding:"required" json:"email"`
	Age       uint8          `binding:"required" json:"age"`
	City      *string        `json:"city"` // When passed * means that this field can be null
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
