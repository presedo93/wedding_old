package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/presedo93/wedding/back/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
