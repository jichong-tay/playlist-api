package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = 3000

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(":%d", port))
}
