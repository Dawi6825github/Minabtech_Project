package models

import (
	"gorm.io/gorm"
)

// User struct with GORM support
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
