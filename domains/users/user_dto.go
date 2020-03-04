package users

import (
	"strings"

	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

// User sctruct describes a user object's properties
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// Validate function validates the user data entries
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalind user email.")
	}
	return nil
}
