package main

import (
	"log"

	"github.com/karlgama/go-microservices/internal/database"
	"github.com/karlgama/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	serv := server.NewEchoServer(db)
	if err := serv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
