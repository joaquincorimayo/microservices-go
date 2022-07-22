package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// custom middleware intercepts and prints the User-Agent header's value for each HTTP request
func FindUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.GetHeader("User-Agent"))
		c.Next()
	}
}

func main () {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(FindUserAgent())
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"Middleware works!"})
	})

	// Start the http server instance
	router.Run(":5000")
}
