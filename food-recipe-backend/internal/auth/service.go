package auth

import (
	"food-recipe-backend/internal/models"
	"food-recipe-backend/internal/users"
	"food-recipe-backend/internal/utils"
)

// SignupService processes signup logic and user creation
func SignupService(signupData *models.SignupRequest) (string, error) {
	hashedPassword, err := utils.HashPassword(signupData.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    signupData.Email,
		Username: signupData.Username,
		Password: hashedPassword,
	}

	if err := users.CreateUser(&user); err != nil {
		return "", err
	}

	return GenerateJWT(user.ID, user.Username)
}

// LoginService processes login logic and returns JWT token
func LoginService(loginData *models.LoginRequest) (string, error) {
	user, err := users.GetUserByEmail(loginData.Email)
	if err != nil || !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return "", err
	}

	return GenerateJWT(user.ID, user.Username)
}
