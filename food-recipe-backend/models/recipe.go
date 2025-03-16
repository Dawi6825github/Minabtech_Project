package models

import (
	"time"

	"gorm.io/gorm"
)

// Recipe represents a recipe in the system
type Recipe struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	CategoryID   uint      `gorm:"not null" json:"category_id"`
	Title        string    `gorm:"not null" json:"title"`
	Description  string    `gorm:"not null" json:"description"`
	Ingredients  string    `gorm:"not null" json:"ingredients"`
	Instructions string    `gorm:"not null" json:"instructions"`
	ImageURL     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

// CreateRecipe creates a new recipe
func CreateRecipe(db *gorm.DB, recipe *Recipe) error {
	return db.Create(recipe).Error
}

// GetRecipeByID fetches a recipe by its ID
func GetRecipeByID(db *gorm.DB, id uint) (*Recipe, error) {
	var recipe Recipe
	err := db.Where("id = ?", id).First(&recipe).Error
	return &recipe, err
}

// GetRecipesByCategory fetches recipes by category
func GetRecipesByCategory(db *gorm.DB, categoryID uint) ([]Recipe, error) {
	var recipes []Recipe
	err := db.Where("category_id = ?", categoryID).Find(&recipes).Error
	return recipes, err
}
