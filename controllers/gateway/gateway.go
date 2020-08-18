package gateway

import (
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/services/routes_service"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
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
		case "/cannabis":
			jwt := c.Request.Header.Get("Authorization")
			r, err := routes_service.RoutesService.CallCannabis("POST", "http://172.30.0.3:8083/api/new-cannabis", c.Request.Body, jwt)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}

			c.JSON(http.StatusCreated, r)
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
