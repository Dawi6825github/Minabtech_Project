package models

import (
	"time"
	"gorm.io/gorm"
)

// Rating represents a rating for a recipe
type Rating struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RecipeID  uint      `gorm:"not null" json:"recipe_id"`
	Score     float32   `gorm:"not null" json:"score"`  // Rating between 0 and 5
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    User    `gorm:"foreignKey:UserID"`
	Recipe  Recipe  `gorm:"foreignKey:RecipeID"`
}

// CreateRating creates a new rating
func CreateRating(db *gorm.DB, rating *Rating) error {
	return db.Create(rating).Error
}

// GetRatingsByRecipe fetches ratings for a specific recipe
func GetRatingsByRecipe(db *gorm.DB, recipeID uint) ([]Rating, error) {
	var ratings []Rating
	err := db.Where("recipe_id = ?", recipeID).Find(&ratings).Error
	return ratings, err
}
