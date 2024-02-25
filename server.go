package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func startServer() {
	r := gin.Default()

	// Define routes
	r.GET("/", handler)
	r.GET("/latest", latestHandler)
	r.POST("/convert", convertHandler)

	// Start the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
