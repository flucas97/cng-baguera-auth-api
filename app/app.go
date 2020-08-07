package app

import (
	"sync"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/gateway"
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/middlewares"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
	Waitg  sync.WaitGroup
)

func StartApp() {
	// Auth Middleware
	Router.Use(middlewares.Entry)
	// Entrypoint API
	Router.Use(gateway.Entry)

	logger.Info("Starting server...")
	Router.Run(":8082")
}
