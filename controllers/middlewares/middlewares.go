package middlewares

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	var (
		reqToken = c.Request.Header["Authorization"]
	)

	if reqToken != nil {
		givenToken := reqToken[0]
		claims := jwt.MapClaims{}

		jwtToken, err := auth.GetJWT(givenToken, claims)
		if err != nil {
			logger.MiddlewareError(err.Error())
		}

		for key, val := range claims {
			fmt.Printf("Key: %v, value: %v\n", key, val)
		}
		_ = jwtToken
		logger.MiddlewareInfo(fmt.Sprintf("protect path %v", claims["name"]))
	} else {
		logger.MiddlewareInfo(fmt.Sprintf("attempt to enter from IP %v", c.Request.Host))
	}
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
