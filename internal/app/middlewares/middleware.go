package middlewares

import (
	timeouts "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return timeouts.New(timeouts.WithTimeout(timeout), timeouts.WithHandler(func(context *gin.Context) {
		context.JSON(http.StatusGatewayTimeout, gin.H{"message": "Request Timeout"})
		context.Abort()
	}))
}
