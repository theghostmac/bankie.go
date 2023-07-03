package users

import "database/sql"

func ScanIntoAccount(rows *sql.Rows) (*CustomerAccount, error) {
	account := new(CustomerAccount)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.BankNumber,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
