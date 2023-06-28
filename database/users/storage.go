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

func (us *UserRepository) Init() error {
	return us.createAccountTable()
}

func (us *UserRepository) createAccountTable() error {
	query := `Create table account if it doesn't exists (
		id serial primary_key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance,
		created_at timestamp
		)`
	_, err := us.db.Exec(query)
	return err
}

func (us *UserRepository) CreateAccount(*CustomerAccount) error {
	return nil
}

func (us *UserRepository) GetAccountByID(id int) *CustomerAccount {
	return nil
}

func (us *UserRepository) GetAccountByEmail(email string) *CustomerAccount {
	return nil
}

func (us *UserRepository) UpdateAccount(*CustomerAccount) error {
	return nil
}

func (us *UserRepository) DeleteAccount(id int) error {
	return nil
}
