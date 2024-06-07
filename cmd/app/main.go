package main

import (
	"cmd/app/main.go/internal/app/handlers"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	server := gin.Default()
	routes := server.Group("/v1")

	routes.Use(gin.LoggerWithWriter(log.Writer()))
	routes.Use(limit.MaxAllowed(5000))

	handlers.NewRouteHandler().Routes(routes)

	log.Info("HTTP ", "addr ", "8081")
	log.Info("Starting Application at: ", server.Run(":8081"))
}
