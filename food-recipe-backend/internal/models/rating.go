package models

import (
	"github.com/google/uuid"

)

type Rating struct {
	Base
	UserID   uuid.UUID
	RecipeID uuid.UUID
	Score    int `gorm:"not null;check:score >= 1 AND score <= 5"`

	User   User   `gorm:"foreignKey:UserID"`
	Recipe Recipe `gorm:"foreignKey:RecipeID"`
}
