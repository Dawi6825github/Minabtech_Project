package recipes

import (
	"fmt"
	// "github.com/dawit6825/food-recipe-backend/internal/db"
	"github.com/jinzhu/gorm"
)

// RecipeService contains methods to interact with the database
type RecipeService struct {
	DB *gorm.DB
}

// NewRecipeService creates a new RecipeService instance
func NewRecipeService(db *gorm.DB) *RecipeService {
	return &RecipeService{DB: db}
}

// GetAll retrieves all recipes from the database
func (s *RecipeService) GetAll() ([]Recipe, error) {
	var recipes []Recipe
	if err := s.DB.Preload("Ingredients").Preload("Steps").Find(&recipes).Error; err != nil {
		return nil, err
	}
	return recipes, nil
}

// Create creates a new recipe in the database
func (s *RecipeService) Create(recipe *Recipe) (*Recipe, error) {
	if err := s.DB.Create(recipe).Error; err != nil {
		return nil, err
	}

	// After creating the recipe, create its ingredients and steps
	for _, ingredient := range recipe.Ingredients {
		ingredient.RecipeID = recipe.ID
		if err := s.DB.Create(&ingredient).Error; err != nil {
			return nil, err
		}
	}

	for _, step := range recipe.Steps {
		step.RecipeID = recipe.ID
		if err := s.DB.Create(&step).Error; err != nil {
			return nil, err
		}
	}

	return recipe, nil
}

// Update updates an existing recipe in the database
func (s *RecipeService) Update(id uint, updatedRecipe *Recipe) (*Recipe, error) {
	var recipe Recipe
	if err := s.DB.Preload("Ingredients").Preload("Steps").First(&recipe, id).Error; err != nil {
		return nil, err
	}

	// Update the recipe details
	recipe.Title = updatedRecipe.Title
	recipe.Description = updatedRecipe.Description
	recipe.Category = updatedRecipe.Category
	recipe.PrepTime = updatedRecipe.PrepTime

	if err := s.DB.Save(&recipe).Error; err != nil {
		return nil, err
	}

	// Update Ingredients and Steps
	for _, ingredient := range updatedRecipe.Ingredients {
		ingredient.RecipeID = recipe.ID
		if err := s.DB.Save(&ingredient).Error; err != nil {
			return nil, err
		}
	}

	for _, step := range updatedRecipe.Steps {
		step.RecipeID = recipe.ID
		if err := s.DB.Save(&step).Error; err != nil {
			return nil, err
		}
	}

	return &recipe, nil
}

// Delete deletes a recipe by its ID
func (s *RecipeService) Delete(id string) error {
	var recipe Recipe
	if err := s.DB.First(&recipe, id).Error; err != nil {
		return err
	}

	if err := s.DB.Delete(&recipe).Error; err != nil {
		return err
	}

	return nil
}

// Helper function to handle database queries for specific recipes
func (s *RecipeService) GetByID(id uint) (*Recipe, error) {
	var recipe Recipe
	if err := s.DB.Preload("Ingredients").Preload("Steps").First(&recipe, id).Error; err != nil {
		return nil, fmt.Errorf("recipe not found: %v", err)
	}
	return &recipe, nil
}
