package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"vuongtran/learning/simplebank/api"
	db "vuongtran/learning/simplebank/db/sqlc"
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
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
