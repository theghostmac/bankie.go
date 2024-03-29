package migrations

import (
	"fmt"
	"github.com/theghostmac/bankie.go/common/logger"
	"github.com/theghostmac/bankie.go/database/users"

	_ "github.com/lib/pq"
)

func DbConn() {
	Store, err := users.NewUserRepository()
	if err != nil {
		logger.ErrorLogs("Failed to create new store...")
	}
	fmt.Printf("%+v\n", Store)
}
