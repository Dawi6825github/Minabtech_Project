package utils

import (
	"errors"
	"regexp"
)

// IsValidEmail checks if an email format is valid
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidPassword checks if the password meets security requirements
func IsValidPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecialChar := regexp.MustCompile(`[\W]`).MatchString(password)

	if !hasNumber || !hasSpecialChar {
		return errors.New("password must include at least one number and one special character")
	}

	return nil
}
