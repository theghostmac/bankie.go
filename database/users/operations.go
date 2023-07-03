package users

import "database/sql"

func ScanIntoAccount(rows *sql.Rows) (*CustomerAccount, error) {
	account := new(CustomerAccount)
	err := rows.Scan(
		&account.FirstName,
		&account.LastName,
		&account.BankNumber,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
