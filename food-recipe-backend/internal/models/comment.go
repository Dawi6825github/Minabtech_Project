package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	Base
	Content  string `gorm:"type:text;not null"`
	UserID   uuid.UUID
	RecipeID uuid.UUID

	User   User   `gorm:"foreignKey:UserID"`
	Recipe Recipe `gorm:"foreignKey:RecipeID"`
}
