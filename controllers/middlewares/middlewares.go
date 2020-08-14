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
	ctx         = context.Background()
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
		if len(reqToken) != 0 {
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
		ok, err := authService.Validate(reqToken[0], ctx)
		if err != nil || !ok {
			ForbiddenPath(c)
			return
		}

		c.AbortWithStatusJSON(http.StatusFound, "already logged in")
		return
	} else {
		switch c.Request.RequestURI {
		case "/new-account":
			switch c.Request.Method {
			case http.MethodPost:
				w, err := http.Post("http://172.30.0.5:8081/api/new-account", "application/json", c.Request.Body)
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}

				callAuthorize(&ctx, w.Header.Get("nick_name"), "account successfuly created", c)
			}
		case "/login":
			switch c.Request.Method {
			case http.MethodGet:

			case http.MethodPost:
				w, err := http.Post("http://172.30.0.3:8081/api/validate", "application/json", c.Request.Body)
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
					return
				}

				callAuthorize(&ctx, w.Header.Get("nick_name"), "logged in", c)
			}
		}
	}
}

func ForbiddenPath(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s", c.ClientIP()))
	c.AbortWithStatusJSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized"))
}

func callAuthorize(ctx *context.Context, nickName string, finalMessage string, c *gin.Context) {
	jwt, restErr := authService.Authorize(nickName, *ctx)
	if restErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, restErr)
		return
	}

	c.Header("Authorization", jwt)
	c.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"authorization": jwt,
			"message":       finalMessage,
		},
	)
	return
}
