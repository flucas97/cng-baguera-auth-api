package app

import (
	"github.com/flucas97/CNG-checknogreen/baguera-auth/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts server
func StartApp() {
	Routes()
	logger.Info("Starting server...")
	router.Run(":8081")
}
