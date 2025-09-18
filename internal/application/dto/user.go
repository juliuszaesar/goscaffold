package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=1,max=50"`
	LastName  string `json:"last_name" validate:"required,min=1,max=50"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Email     *string `json:"email,omitempty" validate:"omitempty,email"`
	FirstName *string `json:"first_name,omitempty" validate:"omitempty,min=1,max=50"`
	LastName  *string `json:"last_name,omitempty" validate:"omitempty,min=1,max=50"`
}

// UserResponse represents a user in API responses
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListUsersRequest represents the request to list users
type ListUsersRequest struct {
	Limit  int `json:"limit" validate:"min=1,max=100"`
	Offset int `json:"offset" validate:"min=0"`
}

// ListUsersResponse represents the response for listing users
type ListUsersResponse struct {
	Users      []UserResponse `json:"users"`
	Total      int64          `json:"total"`
	Limit      int            `json:"limit"`
	Offset     int            `json:"offset"`
	HasMore    bool           `json:"has_more"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}
