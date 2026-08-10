package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read environment variables (Docker sets these)
	env := os.Getenv("NODE_ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Set Gin mode based on NODE_ENV
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Hello from Go Gin REST API! Running in %s mode.\n", env))
	})

	// Bind to 0.0.0.0:port so Docker can map the port externally
	address := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(address)
}
