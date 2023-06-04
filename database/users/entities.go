package main

type LeadAccount struct {
	ID        int
	FirstName string
	LastName  string
}

type CustomerAccount struct {
	LeadAccount
	BankNumber int64
	Balance    string
}

type NewAccountCreator struct {
}
