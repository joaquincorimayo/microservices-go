package main

import (
	"runtime"
	"github.com/gin-gonic/gin"
)

// I will build two endpoints and handlers

func main () {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router := gin.New()
	// router.Use(gin.Logger())


	// HTTP GET -> RETURN JSON Formatted message
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":"Hello world!",
		})
	})

	// HTTP GET -> RETURN The current OS name in plain text format
	router.GET("/os", func(c *gin.Context){
		c.String(200, runtime.GOOS)
	})

	// Start the http server instance
	router.Run(":5000")
}
