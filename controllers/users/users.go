package users

import (
	"net/http"
	"strconv"
	"users-microservice/domain/users"
	"users-microservice/errors"
	"users-microservice/services"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	userId, paramErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if paramErr != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(200, user)

	return
}

// Create user
func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid request body")
		c.JSON(restError.Status, restError)
		return
	}

	result, createErr := services.CreateUser(user)

	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}
