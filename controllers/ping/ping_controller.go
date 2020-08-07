package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping check application status
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
