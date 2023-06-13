package main

import (
	"github.com/theghostmac/bankie.go/api/rest"
	"github.com/theghostmac/bankie.go/common/logger"
	"github.com/theghostmac/bankie.go/database/migrations"
	"github.com/theghostmac/bankie.go/database/users"
	"log"
)

func main() {
	migrations.DbConn()

	store, err := users.NewUserRepository()
	if err != nil {
		log.Fatal(err)
	}

	serverEngine := rest.NewAPIServer(":8082", store)
	serverEngine.StartServer()
	logger.InfoLogs("Server has started up...\n")
	logger.ErrorLogs("Hello Banker, API Server is running successfully...\n")
}
