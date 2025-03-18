package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password for secure storage in the DB
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ValidatePassword compares a hashed password with a plain-text password
func ValidatePassword(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}
