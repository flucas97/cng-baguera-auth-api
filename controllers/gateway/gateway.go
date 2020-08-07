package gateway

import (
	"net/http"

	"github.com/flucas97/cng/cng-baguera-auth-api/utils/logger"
	"github.com/gin-gonic/gin"
)

func Entry(c *gin.Context) {
	URI := c.Request.RequestURI

	switch c.Request.Method {
	case http.MethodGet:
		switch URI {
		case "/login":
			res, err := http.Post("/api/login", "Auth", c.Request.Body)
			if err != nil {

			}
			_ = res
		case "/account-details":
			// account-details
		case "/cannabis":
			// cannabis
		}
	case http.MethodPost:
		switch URI {
		case "/login":
			// login
		case "/account-details":
			// account-details
		case "/cannabis":
			// cannabis
		}
	case http.MethodPatch:
		switch URI {
		case "/account":
			// account-details
		case "/account-details":
			// cannabis
		}
	default:
		logger.Info("no routes")
	}
}
