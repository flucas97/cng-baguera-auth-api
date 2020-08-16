package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/services/auth_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/success_response"
	"github.com/gin-gonic/gin"
)

const (
	accountsServiceURI = "http://172.30.0.5:8081/api/"
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
			found, err := authService.Validate(ctx, reqToken[0])
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
		ok, err := authService.Validate(ctx, reqToken[0])
		if err != nil || !ok {
			ForbiddenPath(c)
			return
		}

		c.AbortWithStatusJSON(http.StatusFound, success_response.Found("already logged in"))
		return
	} else {
		switch c.Request.RequestURI {
		case "/new-account":
			switch c.Request.Method {
			case http.MethodPost:
				r, err := http.Post(accountsServiceURI+"new-account", "application/json", c.Request.Body)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
					return
				}

				callAuthorize(
					&ctx,
					r.Header.Get("nick_name"),
					r.Header.Get("account_id"),
					r.Header.Get("cannabis_repository_id"),
					"account successfuly created",
					"account already exists",
					c,
				)
				return
			}
		case "/login":
			switch c.Request.Method {
			case http.MethodPost:
				r, err := http.Post(accountsServiceURI+"validate", "application/json", c.Request.Body)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
					return
				}

				callAuthorize(
					&ctx,
					r.Header.Get("nick_name"),
					r.Header.Get("account_id"),
					r.Header.Get("cannabis_repository_id"),
					"successfully login",
					"wrong account or password, try again",
					c,
				)
				return
			}
		}
	}
}

func ForbiddenPath(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s", c.ClientIP()))
	c.AbortWithStatusJSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized"))
	/*
		TODO: clear token from cookie/storage
	*/
}

func callAuthorize(ctx *context.Context, nickName string, accountId string, cannabisRepositoryId string, finalMessage string, errorMessage string, c *gin.Context) {
	if nickName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, error_factory.NewBadRequestError(errorMessage))
		return
	}

	jwt, restErr := authService.Authorize(*ctx, nickName, accountId, cannabisRepositoryId)
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
}
