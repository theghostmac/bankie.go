package migrations

import (
	"github.com/theghostmac/bankie.go/helperFiles"
    "github.com/jinzhu/gorm"
    "github.com/theghostmac/bankie.go/interfaces"
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

func createAccounts() {
    db := connectDB()

    users := &[4]interfaces.User{
        {Username: "MacBobby", Email: "macbobbychibuzor@gmail.com"},
        {Username: "GhostMac", Email: "theghostmac@gmail.com"},
        {Username: "Chibuzor", Email: "peekube01@gmail.com"},
        {Username: "Anna", Email: "alpha.maxrdj@gmail.com"},
    }

    for i := 0; i < len(users); i++ {
        generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
        user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
        db.Create(&user)

        account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
        db.Create(&account)
    }
    defer db.Close()
}

func Migrate() {
    users = &interfaces.User{}
    account = &interfaces.Account{}
    db := connectDB()
    db.AutoMigrate(&User{}, &Account{})
    defer db.Close()
    
    createAccounts()
}

