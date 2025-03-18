package models

import (
)

type Category struct {
	Base
	Name    string    `gorm:"unique;not null"`
	Recipes []Recipe `gorm:"many2many:recipe_categories"`
}
