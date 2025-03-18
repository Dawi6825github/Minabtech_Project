package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base contains common fields for all models
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// BeforeCreate generates UUID before inserting a new record
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
