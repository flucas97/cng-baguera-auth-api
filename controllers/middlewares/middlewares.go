package middlewares

import (
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/gateway"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	gateway.Entry(c)
}

/*
	// middlewares

	request has a jwt in a header or cookie?
	- no -> redirect to /login
	- yes ->
			auth := auth.New(name)
			token := auth.GenerateToken()

			check (GET) against Redis if KEY/VALUE (NAME/TOKEN)
			exists and if it is equal of what we have

			- yes ->
					gateway.Entry(c gin.router) // load routes
			- no ->
				redirect to login and clear this token


*/
