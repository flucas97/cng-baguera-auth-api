package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping check application status
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
