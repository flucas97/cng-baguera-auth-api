package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/services/auth_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	ctx         context.Context
	authService = auth_service.AuthService
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
			found, err := authService.Validate(reqToken[0], ctx)
			if err != nil || !found {
				ForbiddenPath(c)
				return
			}
			c.Next()
			return
		} else {
			ForbiddenPath(c)
			return
		}
	}
}

func allowedPath(reqToken []string, c *gin.Context) {
	if len(reqToken) != 0 {
		c.AbortWithStatus(http.StatusFound)
		return
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

				restErr := authService.Authorize(w.Header.Get("nick_name"), ctx)
				if restErr != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, err)
					return
				}

				//c.Header("Authorization", jwt)
				c.AbortWithStatusJSON(http.StatusCreated, "account successfully created")
				return
			}
		case "/login":
			switch c.Request.Method {
			case http.MethodGet:

			}
		}
	}
}

func ForbiddenPath(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s", c.ClientIP()))
	c.JSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized"))
	c.Abort()
}
