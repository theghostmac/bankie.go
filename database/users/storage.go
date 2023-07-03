package users

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
	connStr := "host=localhost port=5432 user=postgres dbname=bankiestore password=ph03n1x sslmode=disable"
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
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number SERIAL,
		balance NUMERIC,
		created_at TIMESTAMP
	)`
	_, err := us.db.Exec(query)
	return err
}

func (us *UserRepository) CreateAccount(ca *CustomerAccount) error {
	query := `insert into account (
           first_name, last_name, email, bank_number, balance, created_at                 
	)
	values ($1, $2, $3, $4, $5, $6)
    `

	response, err := us.db.Query(
		query,
		ca.FirstName,
		ca.LastName,
		ca.BankNumber,
		ca.Balance,
		ca.CreatedAt,
	)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", response)

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
