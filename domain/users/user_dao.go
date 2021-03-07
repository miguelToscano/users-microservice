package users

import (
	"fmt"
	users_db "users-microservice/databases/mysql"
	"users-microservice/errors"
)

const (
	insertUserQuery     = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	selectUserByIdQuery = "SELECT * FROM users WHERE id = ?"
)

func (user *User) Get() (*User, *errors.RestError) {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	statement, err := users_db.Client.Prepare(selectUserByIdQuery)

	defer statement.Close()

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	selectResult, err := users_db.Client.Query(selectUserByIdQuery, user.Id)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	foundUser := selectResult.Scan()

	return &User{Id: foundUser.id, FirstName: foundUser.first_name, Email: foundUser.email, DateCreated: foundUser.date_created}, nil
}

func (user *User) Save() *errors.RestError {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	statement, err := users_db.Client.Prepare(insertUserQuery)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to write user into the database: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get last user's id from the database: %s", err.Error()))
	}

	user.Id = userId

	defer statement.Close()

	return nil
}
