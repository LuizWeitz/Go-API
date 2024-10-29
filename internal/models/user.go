package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v7()" json:"id"`
	Name      string    `binding:"required" json:"name"`
	Email     string    `gorm:"unique" binding:"required" json:"email"`
	Age       uint8     `binding:"required" json:"age"`
	City      *string   `json:"city"` // can be null
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
