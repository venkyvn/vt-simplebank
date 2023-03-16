package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"vuongtran/learning/simplebank/api"
	db "vuongtran/learning/simplebank/db/sqlc"
	"vuongtran/learning/simplebank/gapi"
	simplebank "vuongtran/learning/simplebank/pb"
	"vuongtran/learning/simplebank/util"
)

//const (
//	dbDriver      = "postgres"
//	dbSource      = "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable"
//	serverAddress = "0.0.0.0:8080"
//)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	grpcServer := grpc.NewServer()
	simplebank.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer) // important -> self document to help internal server can find each other

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create gRPC server", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
