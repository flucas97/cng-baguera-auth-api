package gateway

import (
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	URI := c.Request.RequestURI

	switch c.Request.Method {
	case http.MethodGet:
		switch URI {
		case "/login":
			//res, err := http.Post("/api/login", "Auth", c.Request.Body)
			//if err != nil { }
			//_ = res
		case "/ping":
			ping.Ping(c)
			return
		default:
			pathNotFound(c)
			return
		}
	case http.MethodPost:
		switch URI {
		case "/login":
			// login
		case "/account-details":
			// account-details
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
}
