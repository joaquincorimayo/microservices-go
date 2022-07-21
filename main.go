package main

import (
	"github.com/gin-gonic/gin"
)

// Generic endpoint handler
func endpointHandler(c *gin.Context) {
	c.String(200, "%s %s", c.Request.Method, c.Request.URL.Path)
}

func main () {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router := gin.New()
	// router.Use(gin.Logger())

	router.GET("/products", endpointHandler)
	router.GET("/products/:productId", endpointHandler)
	router.POST("/products",endpointHandler)
	router.PUT("/products/:productId", endpointHandler)
	router.DELETE("/products/:productId", endpointHandler)

	// Start the http server instance
	router.Run(":5000")
}
