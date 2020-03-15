package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_users-api/domains/services"
	"github.com/shchaslyvyi/bookstore_users-api/domains/users"
	"github.com/shchaslyvyi/bookstore_users-api/utils/errors"
)

// getUserID is the function to Get the user entity
func getUserID(paramUserID string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(paramUserID, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("Invalid user ID - should be a number")
	}
	return userID, nil
}

// Get is the function to Get the user entity
func Get(c *gin.Context) {
	userID, errID := getUserID(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create is a function to Create a user entity
func Create(c *gin.Context) {
	var user users.User
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

// Update is the function to Update the existing user in the DB
func Update(c *gin.Context) {
	userID, errID := getUserID(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Delete is the function to Delete the existing user from the DB
func Delete(c *gin.Context) {
	userID, errID := getUserID(c.Param("user_id"))
	if errID != nil {
		c.JSON(errID.Status, errID)
		return
	}
	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
