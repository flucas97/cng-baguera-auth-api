package gateway

import (
	"fmt"
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/middlewares"
	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/utils/error_factory"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	URI := c.Request.RequestURI

	switch c.Request.Method {
	case http.MethodGet:
		switch URI {
		case "/ping":
			// app.Waitg.Add(1)
			ping.Ping(c)
		default:
			pathNotFound(c)
			middlewares.Abort(c)
			return
		}
	case http.MethodPost:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.Abort(c)
			return
		}
	case http.MethodPatch:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.Abort(c)
			return
		}
	case http.MethodPut:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.Abort(c)
			return
		}
	case http.MethodDelete:
		switch URI {
		default:
			pathNotFound(c)
			middlewares.Abort(c)
			return
		}
	default:
		pathNotFound(c)
		return
	}
	// app.Waitg.Wait()
}

func pathNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, error_factory.NewNotFoundError(fmt.Sprintf("path %v not found :(", c.Request.RequestURI)))
}
