package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jichong-tay/foodpanda-playlist-api/internal/db"
	"github.com/jichong-tay/foodpanda-playlist-api/internal/handlers"
)

// setupRoutes sets up the routes for the application.
func setupRoutes(dbClient *db.Client) *gin.Engine {
	r := gin.Default()

	r.POST("/users", handlers.CreateUser(dbClient))

	return r
}
