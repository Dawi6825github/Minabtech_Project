package models

import (
	"time"
	"gorm.io/gorm"
)

// Like represents a like on a recipe
type Like struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RecipeID  uint      `gorm:"not null" json:"recipe_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User   User    `gorm:"foreignKey:UserID"`
	Recipe Recipe  `gorm:"foreignKey:RecipeID"`
}

// CreateLike creates a new like
func CreateLike(db *gorm.DB, like *Like) error {
	return db.Create(like).Error
}

// GetLikesByRecipe fetches likes for a specific recipe
func GetLikesByRecipe(db *gorm.DB, recipeID uint) ([]Like, error) {
	var likes []Like
	err := db.Where("recipe_id = ?", recipeID).Find(&likes).Error
	return likes, err
}
