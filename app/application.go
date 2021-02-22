package app

import (
	"users-microservice/controllers/health"

	"github.com/gin-gonic/gin"
)

func Start() {
	app := gin.Default()

	app.GET("/health", health.Check)

	app.Run()
}
