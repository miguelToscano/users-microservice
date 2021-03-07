package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

type DatabaseConnection struct {
	username string
	password string
	host     string
	name     string
}

func init() {
	var (
		err error
		env map[string]string
	)

	env, err = godotenv.Read(".env")

	if err != nil {
		panic(err)
	}

	databaseConnection := DatabaseConnection{username: env["DATABASE_USERNAME"], password: env["DATABASE_PASSWORD"], host: env["DATABASE_HOST"], name: env["DATABASE_NAME"]}

	databaseSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", databaseConnection.username, databaseConnection.password, databaseConnection.host, databaseConnection.name)

	Client, err = sql.Open("mysql", databaseSource)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database connection succesful")
}
