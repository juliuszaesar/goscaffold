package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/juliuszaesar/goscaffold/internal/domain/entity"
	"github.com/juliuszaesar/goscaffold/internal/domain/repository"
	"github.com/juliuszaesar/goscaffold/internal/infrastructure/database"
)

// userRepository implements the UserRepository interface
type userRepository struct {
	db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) repository.UserRepository {
	return &userRepository{db: db}
}

// Create saves a new user to the database
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (id, email, first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID(),
		user.Email().Value(),
		user.Name().FirstName(),
		user.Name().LastName(),
		user.CreatedAt(),
		user.UpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GetByID retrieves a user by their ID
func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users WHERE id = $1`

	var userID uuid.UUID
	var email, firstName, lastName string
	var createdAt, updatedAt time.Time

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&userID, &email, &firstName, &lastName, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	user, err := entity.ReconstructUser(userID, email, firstName, lastName, createdAt, updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct user: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by their email address
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users WHERE email = $1`

	var userID uuid.UUID
	var userEmail, firstName, lastName string
	var createdAt, updatedAt time.Time

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&userID, &userEmail, &firstName, &lastName, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	user, err := entity.ReconstructUser(userID, userEmail, firstName, lastName, createdAt, updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to reconstruct user: %w", err)
	}

	return user, nil
}

// Update updates an existing user in the database
func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users 
		SET email = $2, first_name = $3, last_name = $4, updated_at = $5
		WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query,
		user.ID(),
		user.Email().Value(),
		user.Name().FirstName(),
		user.Name().LastName(),
		user.UpdatedAt(),
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// Delete removes a user from the database
func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// List retrieves a paginated list of users
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	query := `
		SELECT id, email, first_name, last_name, created_at, updated_at
		FROM users 
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var userID uuid.UUID
		var email, firstName, lastName string
		var createdAt, updatedAt time.Time

		err := rows.Scan(&userID, &email, &firstName, &lastName, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}

		user, err := entity.ReconstructUser(userID, email, firstName, lastName, createdAt, updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to reconstruct user: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// Count returns the total number of users
func (r *userRepository) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int64
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return count, nil
}
