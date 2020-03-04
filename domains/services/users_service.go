package services

import (
	"github.com/shchaslyvyi/bookstore_users-api/domains/users"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

// CreateUser is a function that implements the business logic of the user creation
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser is a function that implements the business logic of the user fetch
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
