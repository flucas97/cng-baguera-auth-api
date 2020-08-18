package gateway

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/flucas97/cng/cng-baguera-auth-api/controllers/ping"
	"github.com/flucas97/cng/cng-baguera-auth-api/domain/auth"
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
			cRepoId, restErr := auth.GetValueFromJwtKey(c.Request.Header.Get("Authorization"), "cannabis_repository_id")
			if restErr != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, restErr)
				return
			}
			resp := callCannabis(c.Request.Body, cRepoId)

			c.JSON(http.StatusCreated, resp)
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

func callCannabis(body io.ReadCloser, cRepoId string) string {
	req, err := http.NewRequest("POST", "http://172.30.0.3:8083/api/new-cannabis", body)
	if err != nil {
		logger.Error("error creating request", err)
		return ""
	}

	req.Header.Set("cannabis_repository_id", cRepoId)
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("error making request", err)
		return ""
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()

	return newStr
}
