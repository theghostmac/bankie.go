package users

import (
	"database/sql"
	"fmt"
)

type Storage interface {
	CreateAccount(Account *CustomerAccount) error
	DeleteAccount(int) error
	UpdateAccount(Account *CustomerAccount) error
	GetAccountByID(int) (Account *CustomerAccount, err error)
	GetAccounts() ([]*CustomerAccount, error)
	GetAccountByEmail(string) (Account *CustomerAccount)
	WithdrawFromAccount(accountID int, amount float64) error
	GetAccountBalance(id int) (float64, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	connStr := "host=bankie-postgres port=5432 user=postgres dbname=bankiestore password=ph03n1x sslmode=disable"
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	return &UserRepository{
		DB: DB,
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
    	email VARCHAR(50),
		bank_number VARCHAR(50),
		balance FLOAT8,
		created_at TIMESTAMP
	)`
	_, err := us.DB.Exec(query)
	return err
}

func (us *UserRepository) CreateAccount(ca *CustomerAccount) error {
	query := `insert into account (
           first_name, last_name, email, bank_number, balance, created_at                 
	)
	values ($1, $2, $3, $4, $5, $6)
    `

	response, err := us.DB.Query(
		query,
		ca.FirstName,
		ca.LastName,
		ca.BankNumber,
		ca.Email,
		ca.Balance,
		ca.CreatedAt,
	)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", response)

	return nil
}

func (us *UserRepository) GetAccounts() ([]*CustomerAccount, error) {
	rows, err := us.DB.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*CustomerAccount{}

	for rows.Next() {
		account, err := ScanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (us *UserRepository) GetAccountByID(id int) (*CustomerAccount, error) {
	rows, err := us.DB.Query("SELECT * FROM account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (us *UserRepository) GetAccountByEmail(email string) *CustomerAccount {
	return nil
}

func (us *UserRepository) UpdateAccount(*CustomerAccount) error {
	return nil
}

func (us *UserRepository) DeleteAccount(id int) error {
	_, err := us.DB.Query("DELETE FROM account WHERE id = $1")
	return err
}

func (us *UserRepository) WithdrawFromAccount(accountID int, amount float64) error {
	query := "UPDATE account SET balance = balance - $1 WHERE id = $2"
	_, err := us.DB.Exec(query, amount, accountID)
	return err
}

func (us *UserRepository) GetAccountBalance(id int) (float64, error) {
	query := "SELECT balance FROM account WHERE id = $1"
	var balance float64
	err := us.DB.QueryRow(query, id).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
