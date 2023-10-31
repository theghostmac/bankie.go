package main

import (
	retry "github.com/avast/retry-go"
	"github.com/theghostmac/bankie.go/api/rest"
	"github.com/theghostmac/bankie.go/common/logger"
	"github.com/theghostmac/bankie.go/database/migrations"
	"github.com/theghostmac/bankie.go/database/users"
	"log"
	"time"
)

func main() {
	migrations.DbConn()

	// Define a retry strategy
	retryOptions := []retry.Option{
		retry.Attempts(5),
		retry.Delay(1 * time.Second),
	}

	// Wrap the connection establishment in a retry block.
	err := retry.Do(func() error {
		store, err := users.NewUserRepository()
		if err != nil {
			log.Fatal(err)
		}

		if err := store.Init(); err != nil {
			log.Fatal(err)
		}

		serverEngine := rest.NewAPIServer(":8082", store)
		serverEngine.StartServer()
		logger.InfoLogs("Server has started up...\n")
		logger.ErrorLogs("Hello Banker, API Server is running successfully...\n")

		return nil
	}, retryOptions...)

	if err != nil {
		log.Fatalf("Failed to establish a database connection after retries: %v", err)
	}
}
