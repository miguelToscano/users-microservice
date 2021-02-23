package users

import (
	"net/http"
	"users-microservice/domain/users"
	"users-microservice/errors"
	"users-microservice/services"

	"github.com/gin-gonic/gin"
)

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
