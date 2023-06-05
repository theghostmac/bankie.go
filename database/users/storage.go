package users

import (
	"database/sql"
)

type Storage interface {
	CreateAccount(Account *CustomerAccount) error
	DeleteAccount(int) error
	UpdateAccount(Account *CustomerAccount) error
	GetAccountByID(int) (Account *CustomerAccount)
	GetAccountByEmail(string) (Account *CustomerAccount)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	connStr := "user=bankiestore dbname=postgres password=ph03n1x sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &UserRepository{
		db: db,
	}, nil
}
