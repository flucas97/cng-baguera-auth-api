package app

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/gateway"
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
)

// Routes map all avaliable routes
func StartRoutes() {
	router.GET("/ping", ping.Ping)

	router.GET("/login", gateway.Entry)
	// router.POST("/login", account.Login)

}
