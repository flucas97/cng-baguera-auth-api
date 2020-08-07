package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	var (
		reqToken = c.Request.Header["Authorization"]
	)

	switch c.Request.RequestURI {
	case "/login", "/create-account":
		allowedPath(reqToken, c)
	default:
		if reqToken != nil {
			givenToken := reqToken[0]
			claims := jwt.MapClaims{}

			jwtToken, err := auth.GetJWT(givenToken, claims)
			if err != nil {
				logger.MiddlewareError(err.Error())
			}

			nickName, err := auth.ValidateJWT(jwtToken)
			if err != nil {
				Abort(c)
			}

			_ = nickName

			logger.MiddlewareInfo(fmt.Sprintf("protect path %v", claims["name"]))
		} else {
			Abort(c)
		}
	}
}

func allowedPath(reqToken interface{}, c *gin.Context) {
	if reqToken != nil {

	} else {
		switch c.Request.RequestURI {
		case "/create-account":
			switch c.Request.Method {
			case http.MethodPost:
				resp, err := http.Post("localhost:8081/api/create-account", "Authorized", c.Request.Body)
				if err != nil {
					c.JSON(http.StatusBadRequest, err.Error())
				}
				auth.New()
				c.JSON(http.StatusOK, resp.Body)
			}
		}
	}
}

func Abort(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s", c.ClientIP()))
	c.JSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized"))
	c.Abort()
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
