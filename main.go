package main

import (
	_ "github.com/analect-io/teste-loki/config"
	"github.com/analect-io/teste-loki/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app teste-loki",
			"headers": c.Request.Header,
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}
