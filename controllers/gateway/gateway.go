package gateway

import (
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/services/cannabis_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/gin-gonic/gin"
)

var (
	cs = cannabis_service.CannabisService
)

// Entry
func Entry(c *gin.Context) {
	URI := c.Request.RequestURI

	switch c.Request.Method {
	case http.MethodGet:
		switch URI {
		case "/ping":
			ping.Ping(c)
			return
		case "/cannabis":
			jwt := c.Request.Header.Get("Authorization")
			result, err := cs.FindAllCannabis(jwt)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			c.JSON(http.StatusCreated, result)

			return
		default:
			pathNotFound(c)
			return
		}
	case http.MethodPost:
		switch URI {
		case "/cannabis":
			jwt := c.Request.Header.Get("Authorization")
			result, err := cs.New(c.Request.Body, jwt)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			c.JSON(http.StatusCreated, result)
			return
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
