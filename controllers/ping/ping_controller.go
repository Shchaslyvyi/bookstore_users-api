package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a function to performe a Ping request
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
