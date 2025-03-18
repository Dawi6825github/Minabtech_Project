package users

import (
	"food-recipe-backend/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// NewUserService initializes UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// Fix undefined db issues
func (s *UserService) CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
