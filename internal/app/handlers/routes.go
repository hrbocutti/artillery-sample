package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) Routes(rg *gin.RouterGroup) {
	rg.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rg.GET("/full_capacity", fullCapacityHandler)

	rg.GET("/timeouts", timeoutsHandler)

	rg.GET("/slow_response", slowResponseHandler)

	rg.GET("/error_response", errorResponseHandler)

	rg.GET("/normal_response", normalResponseHandler)

}

func fullCapacityHandler(context *gin.Context) {
	context.JSON(http.StatusServiceUnavailable, gin.H{
		"message": "Service unavailable",
	})
}

func timeoutsHandler(context *gin.Context) {
	time.Sleep(2 * time.Second)
	context.JSON(http.StatusGatewayTimeout, gin.H{
		"message": "Timeouts",
	})
}

func slowResponseHandler(context *gin.Context) {
	time.Sleep(2 * time.Second)
	context.JSON(http.StatusOK, gin.H{
		"message": "Slow response",
	})
}

func errorResponseHandler(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error",
	})
}

func normalResponseHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Normal response",
	})
}
