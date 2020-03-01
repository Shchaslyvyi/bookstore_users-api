package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_users-api/domains/services"
	"github.com/shchaslyvyi/bookstore_users-api/domains/users"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

// GetUser is the function to Get the user entity
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented exception")
}

// CreateUser is a function to Create a user entity
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}
