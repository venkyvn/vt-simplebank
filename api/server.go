package api

import (
	"github.com/gin-gonic/gin"
	db "vuongtran/learning/simplebank/db/sqlc"
)

// Server servers HTTP request for banking services.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
