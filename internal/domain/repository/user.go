package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/juliuszaesar/goscaffold/internal/domain/entity"
)

// UserRepository defines the interface for user persistence operations
type UserRepository interface {
	// Create saves a new user to the repository
	Create(ctx context.Context, user *entity.User) error

	// GetByID retrieves a user by their ID
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	// GetByEmail retrieves a user by their email address
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update updates an existing user in the repository
	Update(ctx context.Context, user *entity.User) error

	// Delete removes a user from the repository
	Delete(ctx context.Context, id uuid.UUID) error

	// List retrieves a paginated list of users
	List(ctx context.Context, limit, offset int) ([]*entity.User, error)

	// Count returns the total number of users
	Count(ctx context.Context) (int64, error)
}
