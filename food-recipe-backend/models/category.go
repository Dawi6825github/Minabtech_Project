package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents a category for recipes
type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateCategory creates a new category
func CreateCategory(db *gorm.DB, category *Category) error {
	return db.Create(category).Error
}

// GetCategories fetches all categories
func GetCategories(db *gorm.DB) ([]Category, error) {
	var categories []Category
	err := db.Find(&categories).Error
	return categories, err
}
