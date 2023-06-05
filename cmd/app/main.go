package main

import (
	"github.com/theghostmac/bankie.go/api/rest"
	"github.com/theghostmac/bankie.go/common/logger"
	"github.com/theghostmac/bankie.go/database/migrations"
)

func main() {
	migrations.DbConn()

	serverEngine := rest.NewAPIServer(":8082", nil)
	serverEngine.StartServer()
	logger.InfoLogs("Server has started up...\n")
	logger.ErrorLogs("Hello Banker, API Server is running successfully...\n")
}
