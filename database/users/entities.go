package users

import (
	"context"
	"time"
)

// custom types.
type (
	AName  string
	Money  string
	BankID int64
)

// TypeOfUser uses number to tell the type of user incoming to bankie.
type TypeOfUser = int

const (
	unknownUserType TypeOfUser = iota
	lead
	customer
	churned
	lostLead
)

// PaymentDetails records the payment details from the external api for funding bankie wallet.
type PaymentDetails struct {
	APISourceTokenID string
}

// UserAddRequest makes any TypeOfUser to become a Customer.
type UserAddRequest struct {
	TypeOfUser     TypeOfUser
	Email          string
	PaymentDetails PaymentDetails
}

type LeadAccount struct {
	Email     string `json:"email"`
	FirstName AName  `json:"firstName"`
	LastName  AName  `json:"lastName"`
}

type CustomerAccount struct {
	LeadAccount
	BankNumber BankID    `json:"bankNumber"`
	Balance    Money     `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
}

type NewAccountCreator struct {
}

// TODO implement these as ubiquitous language for this project:
//  - everything below this line.

// LeadRequest requests for email from a user to become a Lead. See domain file for definition of Lead.
type LeadRequest struct {
	Email string
}

// Lead struct defines ID for a new lead.
type Lead struct {
	ID string
}

// LeadCreator is a collection of methods that makes a user a lead.
type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
}

// Customer struct defines IDs for a new Customer.
type Customer struct {
	leadID string
	userID string
}

// UserID provides an ID for a user.
func (c *Customer) UserID() string {
	return c.userID
}

// SetUserID rephrases UserID for future clarity.
func (c *Customer) SetUserID(userID string) {
	c.userID = userID
}

// LeadConvertor is a collection of methods that converts a Lead into a Customer.
type LeadConvertor interface {
	Convert(ctx context.Context) (Customer, error)
}

// Convert function converts Lead to Customer.
func (l Lead) Convert(ctx context.Context) (Customer, error) {
	// TODO implement converter logic.
	panic("implement converter logic!")
}
