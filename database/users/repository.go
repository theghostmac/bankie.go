package users

import (
	"time"
)

type CreateCustomerRequest struct {
	FirstName AName  `json:"firstName"`
	LastName  AName  `json:"lastName"`
	Email     string `json:"email"`
}

// NewCustomer creates a new customer for bankie.
func NewCustomer(firstName, lastName AName, email string) *CustomerAccount {
	return &CustomerAccount{
		LeadAccount: LeadAccount{
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		},
		//BankNumber: BankID(int64(rand.Intn(1000))),
		Balance:   Money("$300,000,000"),
		CreatedAt: time.Now().UTC(),
	}
}
