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
	case "/login", "/new-account":
		allowedPath(reqToken, c)
		return
	default:
		if reqToken != nil {
			givenToken := reqToken[0]
			claims := jwt.MapClaims{}

			jwtToken, err := auth.ValidateJWT(givenToken, claims)
			if err != nil {
				logger.MiddlewareError(err.Error())
				return
			}

			_ = jwtToken
			logger.MiddlewareInfo(fmt.Sprintf("protect path %v", claims["name"]))
			return
		} else {
			ForbiddenPath(c)
			return
		}
	}
}

func allowedPath(reqToken []string, c *gin.Context) {
	if len(reqToken) != 0 {

	} else {
		switch c.Request.RequestURI {
		case "/new-account":
			switch c.Request.Method {
			case http.MethodPost:
				w, err := http.Post("http://localhost:8081/api/new-account", "Authorized", c.Request.Body)
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}

				nickName, found := getAccountIdentifier(w)
				if !found {
					c.AbortWithStatusJSON(http.StatusBadRequest, error_factory.NewBadRequestError("user already exists"))
					return
				}

				auth := auth.New(nickName)

				jwt, restErr := auth.GenerateJWT()
				if restErr != nil {
					c.JSON(http.StatusInternalServerError, "error generating token, try again")
					return
				}

				/*
					save nickName(key) - jwt(value) into Redis

				*/
				c.Header("cngAuth", jwt)
				c.AbortWithStatusJSON(http.StatusCreated, "account successfully created")
				return
			}
		}
	}
}

func ForbiddenPath(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s", c.ClientIP()))
	c.JSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized"))
	c.Abort()
}

func getAccountIdentifier(w *http.Response) (string, bool) {
	nickName := w.Header.Get("nick_name")
	if nickName == "" {
		return "", false
	}

	return nickName, true
}
