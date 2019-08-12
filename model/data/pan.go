package model

import (
	"encoding/json"
	"time"
)

// PAN is a Key type that encapsulates PAN data.
//
type PAN struct {
	id             string
	cardNumber     string
	cardHolder     string
	experationDate time.Time
	billingAddress string
}

// panJSON is a struct that facilitiates Unmarshaling data from JSON to a pan struct.
//
type panJSON struct {
	ID             string    `json:"id"`
	CardNumber     string    `json:"cardNumber"`
	CardHolder     string    `json:"cardHolder"`
	ExperationDate time.Time `json:"experationDate"`
	BillingAddress string    `json:"billingAddress"`
}

// NewPAN returns a pointer to a new PAN struct.
//
func NewPAN() *PAN {
	return new(PAN)
}

// NewPANFromJSON takes a JSON string and returns a pointer to a new PAN struct with vales from the JSON.
//
func NewPANFromJSON(newPan string) *PAN {
	jsonPan := new(panJSON)
	json.Unmarshal([]byte(newPan), jsonPan)
	finalPan := PAN{jsonPan.ID, jsonPan.CardNumber, jsonPan.CardHolder, jsonPan.ExperationDate, jsonPan.BillingAddress}
	return &finalPan
}
// ToJSON converts a pan struct to a JSON string.
//
func (pan PAN) ToJSON() []byte {
	newJSON := panJSON{pan.id, pan.cardNumber, pan.cardHolder, pan.experationDate, pan.billingAddress}
	result, _ := json.Marshal(newJSON)
	return result
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
