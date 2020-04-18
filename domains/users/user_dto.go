package users

import (
	"strings"

	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
)

// Status definitions
const (
	StatusActive = "aktive"
)

// User sctruct describes a user object's properties in Data Transfer Object.
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users is a slice of users
type Users []User

// Validate function validates the user data entries
func (user *User) Validate() *rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Invalind user email.")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("Invalid password.")
	}
	return nil
}
