package app

import "github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"

// Routes map all avaliable routes
func Routes() {
	router.GET("/api/ping", ping.Ping)

}
