package mysql_utils

import (
	"strings"

	"github.com/mysql"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

// ParseError parsec the errors
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("No record is matching a given id.")
		}
		return errors.NewInternalServerError("Error while parsing database responce.")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data.")
	}
	return errors.NewInternalServerError("Error while parsing database responce.")
}
