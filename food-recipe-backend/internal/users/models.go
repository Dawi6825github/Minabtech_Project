package users

import (
    "time"
)

// User represents a user in the system
type User struct {
    ID        int       `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// UserUpdateRequest represents the fields needed to update a user
type UserUpdateRequest struct {
    FirstName *string `json:"first_name,omitempty"`
    LastName  *string `json:"last_name,omitempty"`
    Email     *string `json:"email,omitempty"`
}

// UserCreateRequest represents the fields needed to create a new user
type UserCreateRequest struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}
