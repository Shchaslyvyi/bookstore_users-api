package users

import (
	"fmt"
	"strings"

	"github.com/shchaslyvyi/bookstore_users-api/datasources/mysql/users_db"
	"github.com/shchaslyvyi/bookstore_users-api/logger"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
	"github.com/shchaslyvyi/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	errorNoRows                 = "no rows in result set"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE from users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

// Get is a persistance leyer of GetUser function in Data Access Object.
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare the get user statement.", err)
		return errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to prepare the get user by id.", getErr)
		return errors.NewInternalServerError("Database error.")
	}
	return nil
}

// Save is a persistance leyer of SaveUser function in Data Access Object
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare save user statement.", err)
		return errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("Error when trying to save the user.", saveErr)
		return errors.NewInternalServerError("Database error.")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get last insert id after creating a new user.", err)
		return errors.NewInternalServerError("Database error.")
	}

	user.ID = userID
	return nil
}

// Update is a persistance leyer of UpdateUser function in Data Access Object
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare update user statement.", err)
		return errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("Error when trying to update the user.", err)
		return errors.NewInternalServerError("Database error.")
	}
	return nil
}

// Delete is a persistance leyer of DeleteUser function in Data Access Object
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement.", err)
		return errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("Error when trying to delete the user.", err)
		return errors.NewInternalServerError("Database error.")
	}
	return nil
}

// Search is a persistance leyer of FindByStatus function in Data Access Object
func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare the find users by status statement.", err)
		return nil, errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error when trying to find user by status.", err)
		return nil, errors.NewInternalServerError("Database error.")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan user row into user struct.", err)
			return nil, errors.NewInternalServerError("Database error.")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError((fmt.Sprintf("No users matching status %s.", status)))
	}
	return results, nil
}

// FindByEmailAndPassword is a persistance leyer of LoginUser function in Data Access Object.
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error when trying to prepare the get user by email and password statement.", err)
		return errors.NewInternalServerError("Database error.")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("Invalid user credentials.")
		}
		logger.Error("Error when trying to prepare the get user by email and password.", getErr)
		return errors.NewInternalServerError("Database error.")
	}
	return nil
}
