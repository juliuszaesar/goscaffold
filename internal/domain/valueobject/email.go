package valueobject

import (
	"fmt"
	"regexp"
	"strings"
)

// Email represents an email address value object
type Email struct {
	value string
}

// emailRegex is a simple email validation regex
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NewEmail creates a new Email value object
func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, fmt.Errorf("email cannot be empty")
	}

	// Normalize email (trim whitespace and convert to lowercase)
	normalized := strings.ToLower(strings.TrimSpace(email))

	if !emailRegex.MatchString(normalized) {
		return Email{}, fmt.Errorf("invalid email format: %s", email)
	}

	return Email{value: normalized}, nil
}

// Value returns the email string value
func (e Email) Value() string {
	return e.value
}

// String implements the Stringer interface
func (e Email) String() string {
	return e.value
}

// Equals checks if two emails are equal
func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
