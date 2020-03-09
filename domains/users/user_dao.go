package users

import (
	"fmt"
	"strings"

	"github.com/shchaslyvyi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/shchaslyvyi/bookstore_users-api/utils/date_utils"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// Get is a persistance leyer getUser function in Data Access Object.
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Internal server error happened: %s.", err.Error()))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("Error, user not found %d", user.ID))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while trying to get from the DB the user %d, %s", user.ID, err.Error()))
	}
	return nil
}

// Save is a persistance leyer getUser function in Data Access Object
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Internal server error happened: %s.", err.Error()))
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists.", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save a user: %s.", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error while trying to save a user: %s.", err.Error()))
	}

	user.ID = userID
	return nil
}
