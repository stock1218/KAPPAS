package model

import "time"

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct {
	id             string
	payload        string
	isEncrypted    bool
	cardNumber     string
	cardHolder     string
	experationDate time.Time
}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}

// GetPayload will return the a string representation of the data stored in this Data.
//
func (pan PAN) GetPayload() string {
	return pan.payload
}

// SetPayload will set the payload of the current Data.
//
func (pan *PAN) SetPayload(newPayload string) {
	pan.payload = newPayload
}

// IsEncrypted returns a bool that indicates if the pan payload is encrypted.
//
func (pan PAN) IsEncrypted() bool {
	return pan.isEncrypted
}

// SetEncrypted sets the encrypted state of the pan.
//
func (pan *PAN) SetEncrypted(newState bool) {
	pan.isEncrypted = newState
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
