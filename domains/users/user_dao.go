package users

import (
	"fmt"

	"github.com/shchaslyvyi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/shchaslyvyi/bookstore_users-api/utils/date_utils"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Get is a persistance leyer getUser function
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// Save is a persistance leyer getUser function
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %s already exists", user.FirstName))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.ID] = user
	return nil
}
