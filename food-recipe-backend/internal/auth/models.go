package auth

import (
	// "food-recipe-backend/config"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUserByEmail fetches a user by their email from the database
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func CreateUser(email, password string) (uint64, error) {
	user := User{Email: email, Password: password}
	err := config.DB.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
