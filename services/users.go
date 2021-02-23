package services

import (
	"users-microservice/domain/users"
	"users-microservice/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
