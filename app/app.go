package app

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/middlewares"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	router.Use(middlewares.Entry)
	StartRoutes()

	logger.Info("Starting server...")
	router.Run(":8082")
}
