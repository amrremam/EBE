package models

import (
	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Status      bool      `gorm:"default:false"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
}