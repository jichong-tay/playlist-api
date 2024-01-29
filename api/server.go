package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jichong-tay/foodpanda-playlist-api/db/sqlc"
)

// Server will serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	server.setupRouter()
	return server
}

// setRouter setup routing
func (server *Server) setupRouter() {
	router := gin.Default()

	// add routes to router
	router.POST("/users", server.createUser)
	router.POST("/playlists", server.createPlaylist)
	router.GET("/playlists/:id", server.getPlaylist)
	router.GET("/playlists", server.listPlaylist)
	router.GET("/playlists-current/:id", server.getPlaylistCurrent)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
