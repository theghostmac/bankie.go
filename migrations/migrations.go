package migrations

import (
	"github.com/theghostmac/bankie.go/helperFiles"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
    gorm.Model
    Username string
    Email string
    Password string
}

type Account struct {
    gorm.Model
    Type string
    Name string
    Balance uint
    UserID uint
}

func connectDB() *gorm.DB {
    db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=BankiePostgresApp password=b4nk1eg0 sslmode=disable")
    helpers.HandleErr(err)
    return db 
}

