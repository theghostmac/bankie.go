package helperFiles


import (
    "golang.org/x/crypto/bcrypt"
)

func HandleErr(err error) {
	if err != nil {
			panic(err.Error())
	}
}

