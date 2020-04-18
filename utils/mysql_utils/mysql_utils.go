package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

// ParseError parsec the errors
func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("No record is matching a given id.")
		}
		return rest_errors.NewInternalServerError("Error while parsing database responce.", sqlErr)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("Invalid data.")
	}
	return rest_errors.NewInternalServerError("Error while parsing database responce.", sqlErr)
}
