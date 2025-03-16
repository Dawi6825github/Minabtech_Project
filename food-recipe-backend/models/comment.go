package models

import (
	"time"
	"gorm.io/gorm"
)

// Comment represents a comment on a recipe
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	RecipeID  uint      `gorm:"not null" json:"recipe_id"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    User    `gorm:"foreignKey:UserID"`
	Recipe  Recipe  `gorm:"foreignKey:RecipeID"`
}

// CreateComment creates a new comment
func CreateComment(db *gorm.DB, comment *Comment) error {
	return db.Create(comment).Error
}

// GetCommentsByRecipe fetches comments for a specific recipe
func GetCommentsByRecipe(db *gorm.DB, recipeID uint) ([]Comment, error) {
	var comments []Comment
	err := db.Where("recipe_id = ?", recipeID).Find(&comments).Error
	return comments, err
}
