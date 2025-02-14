package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Status      bool      `gorm:"default:false"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
}


// BeforeCreate sets a UUID before saving a task
func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New()
	return
}