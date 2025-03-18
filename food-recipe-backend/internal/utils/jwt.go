package utils

import "github.com/dgrijalva/jwt-go"

// JWTClaims defines the structure of JWT claims
type JWTClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}
