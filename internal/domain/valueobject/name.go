package valueobject

import (
	"fmt"
	"strings"
)

// Name represents a person's name value object
type Name struct {
	firstName string
	lastName  string
}

// NewName creates a new Name value object
func NewName(firstName, lastName string) (Name, error) {
	firstName = strings.TrimSpace(firstName)
	lastName = strings.TrimSpace(lastName)

	if firstName == "" {
		return Name{}, fmt.Errorf("first name cannot be empty")
	}

	if lastName == "" {
		return Name{}, fmt.Errorf("last name cannot be empty")
	}

	if len(firstName) > 50 {
		return Name{}, fmt.Errorf("first name cannot exceed 50 characters")
	}

	if len(lastName) > 50 {
		return Name{}, fmt.Errorf("last name cannot exceed 50 characters")
	}

	return Name{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

// FirstName returns the first name
func (n Name) FirstName() string {
	return n.firstName
}

// LastName returns the last name
func (n Name) LastName() string {
	return n.lastName
}

// FullName returns the full name
func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastName)
}

// String implements the Stringer interface
func (n Name) String() string {
	return n.FullName()
}

// Equals checks if two names are equal
func (n Name) Equals(other Name) bool {
	return n.firstName == other.firstName && n.lastName == other.lastName
}
