package app

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/gateway"
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/middlewares"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
	//config = cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://google.com"}
	//config.AddAllowOrigins("http://facebook.com")

)

func StartApp() {
	// Auth Middleware
	Router.Use(middlewares.Entry)
	// Entrypoint API
	Router.Use(gateway.Entry)

	logger.Info("Starting server...")
	Router.Run(":8082")
}
