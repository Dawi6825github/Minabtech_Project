package models

import (
	"github.com/google/uuid"

)

type Favorite struct {
	Base
	UserID   uuid.UUID
	RecipeID uuid.UUID

	User   User   `gorm:"foreignKey:UserID"`
	Recipe Recipe `gorm:"foreignKey:RecipeID"`
}
