package gateway

import (
	"fmt"
	"net/http"

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
			return
		case "/cannabis":
			logger.Info("entry cannabis gateway")
			_, err := http.Get("http://172.30.0.5:8081/api/ping")
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "you can go to cannabis :)",
			})
			return
		default:
			pathNotFound(c)
			return
		}
	case http.MethodPost:
		switch URI {
		default:
			pathNotFound(c)
			return
		}
	case http.MethodPatch:
		switch URI {
		default:
			pathNotFound(c)
			return
		}
	case http.MethodPut:
		switch URI {
		default:
			pathNotFound(c)
			return
		}
	case http.MethodDelete:
		switch URI {
		default:
			pathNotFound(c)
			return
		}
	default:
		pathNotFound(c)
		return
	}
}

func pathNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, error_factory.NewNotFoundError(fmt.Sprintf("path %v not found :(", c.Request.RequestURI)))
	c.Abort()
}
