package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "vuongtran/learning/simplebank/db/sqlc"
)

// Server servers HTTP request for banking services.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("currency", validCurrency)
	}
	router.POST("/accounts", server.createAccount)

	router.POST("/users", server.createUser)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
