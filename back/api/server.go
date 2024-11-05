package api

import (
	"github.com/gin-gonic/gin"
	"github.com/presedo93/wedding/back/auth"
	db "github.com/presedo93/wedding/back/db/sqlc"
	"github.com/presedo93/wedding/back/logs"
	"github.com/rs/zerolog/log"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	jwks   auth.JWKS
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store, jwks auth.JWKS) *Server {
	s := &Server{store: store, jwks: jwks}
	r := gin.New()
	r.Use(gin.Recovery()).Use(logs.Middleware())

	api := r.Group("/api").Use(auth.Middleware(s.jwks))

	// Admin routes
	api.GET("/admin/profiles/", s.getProfiles)
	api.GET("/admin/guests", s.getGuests)

	// Profile routes
	api.GET("/profiles", s.getProfile)
	api.POST("/profiles", s.createProfile)
	api.PUT("/profiles", s.updateProfile)
	api.DELETE("/profiles", s.deleteProfile)

	// Profile's guests routes
	api.GET("/profiles/guests", s.getUserGuests)
	api.DELETE("/profiles/guests", s.deleteUserGuests)

	// Guest routes
	api.GET("/guests/:id", s.getGuest)
	api.POST("/guests", s.createGuest)
	api.PUT("/guests/:id", s.updateGuest)
	api.DELETE("/guests/:id", s.deleteGuest)

	s.router = r
	return s
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	log.Error().Err(err).Msg("Server API")
	return gin.H{"error": err.Error()}
}
