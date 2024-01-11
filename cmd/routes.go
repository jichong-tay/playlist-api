package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// setupRoutes sets up the routes for the application.
func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
