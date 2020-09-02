package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/flucas97/cng/cng-baguera-auth-api/services/auth_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/success_response"
	"github.com/gin-gonic/gin"
)

var (
	accountsServiceURI = os.Getenv("ACCOUNTS_BASE_URL")
	ctx                = context.Background()
	authService        = auth_service.AuthService
)

// Entry validate  user informations before authorize to use the rest of the service
func Entry(c *gin.Context) {
	var (
		reqToken = c.Request.Header["Authorization"]
	)

	switch c.Request.RequestURI {
	case "/login", "/signup":
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
		case "/signup":
			switch c.Request.Method {
			case http.MethodPost:
				r, err := http.Post(accountsServiceURI+"new-account", "application/json", c.Request.Body)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
					return
				}
				callAuthorize(
					&ctx,
					r.Header.Get("Nick-name"),
					r.Header.Get("Account-id"),
					r.Header.Get("Repository-id"),
					"account successfuly created.",
					"account already exists.",
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
					r.Header.Get("Nick-name"),
					r.Header.Get("Account-id"),
					r.Header.Get("Repository-id"),
					"successfully login.",
					"wrong account or password, try again.",
					c,
				)
				return
			}
		}
	}
}

// ForbiddenPath blocks the request if its not authorized
func ForbiddenPath(c *gin.Context) {
	logger.MiddlewareAttempt(fmt.Sprintf("attempt to enter from IP %s.", c.ClientIP()))
	c.AbortWithStatusJSON(http.StatusForbidden, error_factory.NewBadRequestError("not authorized."))
	/*
		TODO: clear token from cookie/storage
	*/
}

func callAuthorize(ctx *context.Context, nickName string, accountID string, cannabisRepositoryID string, finalMessage string, errorMessage string, c *gin.Context) {
	if nickName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, error_factory.NewBadRequestError(errorMessage))
		return
	}

	jwt, restErr := authService.Authorize(*ctx, nickName, accountID, cannabisRepositoryID)
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
