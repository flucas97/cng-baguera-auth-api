package gateway

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/middlewares"
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	URI := c.Request.RequestURI

	switch c.Request.Method {
	case http.MethodGet:
		switch URI {
		case "/ping":
			ping.Ping(c)
		case "/cannabis":
			logger.Info("entry cannabis gateway")
			time.Sleep(2 * time.Second)
			c.AbortWithStatusJSON(http.StatusContinue, gin.H{
				"message":     "continue to cannabis :)",
				"status_code": http.StatusContinue,
			})
			return
		default:
			pathNotFound(c)
			middlewares.ForbiddenPath(c)
			return
		}
	case http.MethodPost:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.ForbiddenPath(c)
			return
		}
	case http.MethodPatch:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.ForbiddenPath(c)
			return
		}
	case http.MethodPut:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.ForbiddenPath(c)
			return
		}
	case http.MethodDelete:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.ForbiddenPath(c)
			return
		}
	default:
		pathNotFound(c)
		return
	}
}

func pathNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, error_factory.NewNotFoundError(fmt.Sprintf("path %v not found :(", c.Request.RequestURI)))
}
