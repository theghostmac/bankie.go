package main

import (
	"github.com/theghostmac/bankie.go/api/rest"
)

func main() {
	println("Hello Banker👋!")
	serverEngine := NewApiServer(":8082")
	serverEngine.StartServer()
}
