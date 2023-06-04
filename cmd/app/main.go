package main

import (
	"github.com/theghostmac/bankie.go/api/rest"
	"github.com/theghostmac/bankie.go/common/logger"
)

func main() {
	serverEngine := rest.NewAPIServer(":8082")
	serverEngine.StartServer()
	logger.InfoLogs("Server has started up...\n")
	logger.ErrorLogs("Hello Banker, API Server is running successfully...\n")
}
