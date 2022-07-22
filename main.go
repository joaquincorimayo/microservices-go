package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main () {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"CORS works!"})
	})

	// Start the http server instance
	router.Run(":5000")
}
