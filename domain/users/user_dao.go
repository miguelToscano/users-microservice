package users

import (
	"fmt"
	"users-microservice/errors"
)

// user_dao stands for "User Data access object", in this file we define how a
// user interacts with the database

// mock database
var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDb[user.ID]

	if result == nil {
		err := errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", user.ID))
		return err
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {

	if usersDb[user.ID] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("User with ID: %d already exists", user.ID))
	}

	usersDb[user.ID] = user

	return nil
}
