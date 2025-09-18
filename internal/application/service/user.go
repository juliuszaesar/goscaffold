// Package service contains application services that orchestrate business operations.
// Services coordinate between domain entities and infrastructure components,
// implementing use cases while maintaining clean architecture principles.
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/juliuszaesar/goscaffold/internal/application/dto"
	"github.com/juliuszaesar/goscaffold/internal/domain/entity"
	"github.com/juliuszaesar/goscaffold/internal/domain/repository"
)

// UserService handles user-related business logic
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if user with email already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}

	// Create new user entity
	user, err := entity.NewUser(req.Email, req.FirstName, req.LastName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Save user to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return s.entityToResponse(user), nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return s.entityToResponse(user), nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Update email if provided
	if req.Email != nil {
		if err := user.UpdateEmail(*req.Email); err != nil {
			return nil, fmt.Errorf("failed to update email: %w", err)
		}
	}

	// Update name if provided
	if req.FirstName != nil || req.LastName != nil {
		firstName := user.Name().FirstName()
		lastName := user.Name().LastName()

		if req.FirstName != nil {
			firstName = *req.FirstName
		}
		if req.LastName != nil {
			lastName = *req.LastName
		}

		if err := user.UpdateName(firstName, lastName); err != nil {
			return nil, fmt.Errorf("failed to update name: %w", err)
		}
	}

	// Save updated user
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return s.entityToResponse(user), nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// ListUsers retrieves a paginated list of users
func (s *UserService) ListUsers(ctx context.Context, req dto.ListUsersRequest) (*dto.ListUsersResponse, error) {
	// Set default values
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	users, err := s.userRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	total, err := s.userRepo.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to count users: %w", err)
	}

	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = *s.entityToResponse(user)
	}

	hasMore := int64(req.Offset+len(users)) < total

	return &dto.ListUsersResponse{
		Users:   userResponses,
		Total:   total,
		Limit:   req.Limit,
		Offset:  req.Offset,
		HasMore: hasMore,
	}, nil
}

// entityToResponse converts a user entity to a response DTO
func (s *UserService) entityToResponse(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID(),
		Email:     user.Email().Value(),
		FirstName: user.Name().FirstName(),
		LastName:  user.Name().LastName(),
		FullName:  user.Name().FullName(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}
