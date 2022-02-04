package users

import (
    "time"

	"github.com/bankie.go/helperFiles"
	"github.com/bankie.go/interfaces"
	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) map[string]interface{} {
	user := &interfaces.User{}
	
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
    return map[string]interface{}{"message": "User not found"}
}

}