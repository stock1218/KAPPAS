package model

import "time"

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct {
	id             string
	cardNumber     string
	cardHolder     string
	experationDate time.Time
	billingAddress string
}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}

// GetID returns the ID of the pan.
//
func (pan PAN) GetID() string {
	return pan.id
}

// SetID sets the ID of the pan.
//
func (pan *PAN) SetID(newID string) {
	pan.id = newID
}

// GetCardNumber returns the card number of the pan.
//
func (pan PAN) GetCardNumber() string {
	return pan.cardNumber
}

// SetCardNumber sets the card number of the pan.
//
func (pan *PAN) SetCardNumber(newCardNumber string) {
	pan.cardNumber = newCardNumber
}

// GetCardHolder returns the card holder of the pan.
//
func (pan PAN) GetCardHolder() string {
	return pan.cardHolder
}

// SetCardHolder sets the card holder of the pan.
//
func (pan *PAN) SetCardHolder(newCardHolder string) {
	pan.cardHolder = newCardHolder
}

// GetExperationDate returns the experation date of the pan.
//
func (pan PAN) GetExperationDate() time.Time {
	return pan.experationDate
}

// SetExperationDate sets the experation date of the pan.
//
func (pan *PAN) SetExperationDate(newDate time.Time) {
	pan.experationDate = newDate
}

// GetBillingAddress returns the billing address of the pan.
//
func (pan PAN) GetBillingAddress() string {
	return pan.billingAddress
}

// SetBillingAddress sets the billing address of the pan.
//
func (pan *PAN) SetBillingAddress(newAddress string) {
	pan.billingAddress = newAddress
}
