package services

import (
	"github.com/shchaslyvyi/bookstore_users-api/domains/users"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

// CreateUser is a function that implements the business logic of the user creation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
