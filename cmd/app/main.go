package main

import (
	"github.com/theghostmac/bankie.go/api/rest"

	"log"
)

func main() {
	serverEngine := rest.NewAPIServer(":8082")
	serverEngine.StartServer()

	log.Println("Hello Banker, API Server is running successfully...")
}
