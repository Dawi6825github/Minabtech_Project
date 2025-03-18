package auth

import (
	"fmt"
	"net/http"
	"strings"
	"food-recipe-backend/internal/utils"
)

// AuthMiddleware checks for valid JWT token in the request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the authorization header is set
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header (Bearer <token>)
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse and validate the token
		claims, err := ParseJWT(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		// Set the user ID in the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, utils.UserIDKey, claims.UserID)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
