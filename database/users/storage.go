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
	connStr := "user=bankiestore dbname=postgres password=ph03n1x sslmode=disable"
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

func (us *UserRepository) CreateAccount(account *CustomerAccount) error {
	return nil
}

func (us *UserRepository) GetAccountByID(id int) error {
	return nil
}

func (us *UserRepository) GetAccountByEmail(email string) error {
	return nil
}

func (us *UserRepository) UpdateAccount(account *CustomerAccount) error {
	return nil
}

func (us *UserRepository) DeleteAccount(id int) error {
	return nil
}
