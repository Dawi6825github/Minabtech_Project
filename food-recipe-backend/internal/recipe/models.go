package recipes

import "github.com/jinzhu/gorm"

// Recipe model structure
type Recipe struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	PrepTime    int    `json:"prep_time"` // in minutes
	Ingredients []Ingredient `json:"ingredients" gorm:"foreignkey:RecipeID"`
	Steps       []Step `json:"steps" gorm:"foreignkey:RecipeID"`
	CreatedBy   uint   `json:"created_by"`
	Images      []string `json:"images"` // store URLs of images
}

// Ingredient model for dynamic ingredients list
type Ingredient struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	RecipeID  uint   `json:"recipe_id"`
	Name      string `json:"name" binding:"required"`
	Quantity  string `json:"quantity" binding:"required"`
}

// Step model for dynamic steps list
type Step struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	RecipeID  uint   `json:"recipe_id"`
	StepOrder int    `json:"step_order"`
	Instruction string `json:"instruction" binding:"required"`
}
