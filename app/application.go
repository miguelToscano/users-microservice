package app

import (
	"users-microservice/controllers/health"
	"users-microservice/controllers/users"

	"github.com/gin-gonic/gin"
)

func Start() {
	app := gin.Default()

	app.GET("/health", health.Check)
	app.POST("/users", users.Create)
	app.GET("/users/:user_id", users.Get)

	app.Run()
}
