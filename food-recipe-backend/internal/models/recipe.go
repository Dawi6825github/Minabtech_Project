package models

import "github.com/google/uuid"


type Recipe struct {
	Base
	Title       string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	ImageURL    string  
	PrepTime    int     // in minutes
	CookTime    int     // in minutes
	Servings    int     `gorm:"default:1"`
	UserID      uuid.UUID // Foreign key to User

	// Relations
	User       User      `gorm:"foreignKey:UserID"`
	Comments   []Comment `gorm:"foreignKey:RecipeID"`
	Ratings    []Rating  `gorm:"foreignKey:RecipeID"`
	Categories []Category `gorm:"many2many:recipe_categories"`
	Favorites  []Favorite `gorm:"foreignKey:RecipeID"`
}
