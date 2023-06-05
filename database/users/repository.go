package users

import "math/rand"

// NewCustomer creates a new customer for bankie.
func NewCustomer(firstName, lastName AName, email string) *CustomerAccount {
	return &CustomerAccount{
		LeadAccount: LeadAccount{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		},
		BankNumber: BankID(int64(rand.Intn(1000))),
		// FIXME:   ^ this conversion is redundant.
		Balance: Money("$0000"),
	}
}
