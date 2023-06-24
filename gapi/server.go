package gapi

import (
	"fmt"
	db "vuongtran/learning/simplebank/db/sqlc"
	simplebank "vuongtran/learning/simplebank/pb"
	"vuongtran/learning/simplebank/token"
	"vuongtran/learning/simplebank/util"
)

// Server servers gRPC request for banking services.
type Server struct {
	simplebank.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
