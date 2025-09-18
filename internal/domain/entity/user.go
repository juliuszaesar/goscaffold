package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/juliuszaesar/goscaffold/internal/domain/valueobject"
)

// User represents a user in the domain
type User struct {
	id        uuid.UUID
	email     valueobject.Email
	name      valueobject.Name
	createdAt time.Time
	updatedAt time.Time
}

// NewUser creates a new user with the given email and name
func NewUser(email, firstName, lastName string) (*User, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	nameVO, err := valueobject.NewName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		id:        uuid.New(),
		email:     emailVO,
		name:      nameVO,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructUser reconstructs a user from persistence (used by repositories)
func ReconstructUser(id uuid.UUID, email, firstName, lastName string, createdAt, updatedAt time.Time) (*User, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	nameVO, err := valueobject.NewName(firstName, lastName)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        id,
		email:     emailVO,
		name:      nameVO,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

// ID returns the user's ID
func (u *User) ID() uuid.UUID {
	return u.id
}

// Email returns the user's email
func (u *User) Email() valueobject.Email {
	return u.email
}

// Name returns the user's name
func (u *User) Name() valueobject.Name {
	return u.name
}

// CreatedAt returns when the user was created
func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

// UpdatedAt returns when the user was last updated
func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

// UpdateEmail updates the user's email
func (u *User) UpdateEmail(email string) error {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return err
	}
	
	u.email = emailVO
	u.updatedAt = time.Now()
	return nil
}

// UpdateName updates the user's name
func (u *User) UpdateName(firstName, lastName string) error {
	nameVO, err := valueobject.NewName(firstName, lastName)
	if err != nil {
		return err
	}
	
	u.name = nameVO
	u.updatedAt = time.Now()
	return nil
}
