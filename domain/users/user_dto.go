package users

// user_dto stands for "User DAta Transfer Object", in this file we define how a
// user object behaves

import (
	"strings"
	"users-microservice/errors"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
