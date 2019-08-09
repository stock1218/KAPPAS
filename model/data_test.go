package model

import (
	"testing"
	"time"
)

// TestNewData will test if a data variable can be created.
//
func TestNewData(t *testing.T) {
	var _ Data = nil
}

func TestNewPANModel(t *testing.T) {
	var pan Data = NewPAN()
	if pan == nil {
		t.Log("Data: Failed to create NewPAN")
		t.Fail()
	}
}

// PanSetUp is a setup function for creating a testable PAN.
//
func PANSetUp() Data {
	var pan Data = NewPAN()
	return pan
}

// TestSetAndGetPayload tests if SetPayload will set a Data struct's payload GetPayload will return it.
//
func TestSetAndGetPayload(t *testing.T) {
	pan := PANSetUp()
	myPayload := "payload"
	pan.SetPayload(myPayload)
	testPayload := pan.GetPayload()

	if testPayload == "" {
		t.Log("Data: GetPayload returned empty string")
		t.Fail()
	} else if myPayload == testPayload {
		t.Log("Data: GetPayload didn't return correct payload, or SetPayload didn't set correct payload")
	}
}

// TestSetAndGetID tests if SetId will set the id of a given pan struct, and GetId will return it.
//
func TestSetAndGetID(t *testing.T) {
	pan := PANSetUp()
	initialID := "abc123"
	pan.SetID(initialID)

	returnedID := pan.GetID()

	if returnedID == "" {
		t.Log("Data: GetId returned empty string")
		t.Fail()
	} else if initialID != returnedID {
		t.Log("Data: SetId didn't set the correct ID, or GetId returned the wrong ID")
	}
}

// TestSetAndGetCardNumber tests if SetCardNumber will set card-number in a pan struct, and GetCardNumber will return it.
func TestSetAndGetCardNumber(t *testing.T) {
	pan := PANSetUp()
	initialNumber := "123456789"
	pan.(*PAN).SetCardNumber(initialNumber)

	returnedCardNumber := pan.(*PAN).GetCardNumber()
	if returnedCardNumber != initialNumber {
		t.Log("Data: SetCardNumber set the wrong value, or GetCardNumber returned the wrong value")
		t.Fail()
	}
}

// TestSetAndGetCardHolder tests if SetCardHolder will set card-holder in a pan struct, and GetCardHolder will return it.
func TestSetAndGetCardHolder(t *testing.T) {
	pan := PANSetUp()
	initialName := "John Doe"
	pan.(*PAN).SetCardHolder(initialName)

	returnedName := pan.(*PAN).GetCardHolder()
	if initialName != returnedName {
		t.Log("Data: SetCardHolder set the wrong value, or GetCardHolder returned the wrong value")
		t.Fail()
	}

}

// TestSetAndGetExperationDate tests if SetExperationDate will set the experation date in a pan struct,
// and GetCardHolder will return it.
func TestSetAndGetExperationDate(t *testing.T) {
	pan := PANSetUp()
	initialDate := time.Now()

	pan.(*PAN).SetExperationDate(initialDate)
	returnedDate := pan.(*PAN).GetExperationDate()

	if initialDate != returnedDate {
		t.Log("Data: SetExperationDate set the wrong value, or GetExperationDate returned the wrong value")
	}

}

func TestSetAndGetBillingAddress(t *testing.T) {
	pan := PANSetUp()
	initialAddress := "123 4th Ave, Apt. 5E, New York NY 67890"

	pan.(*PAN).SetBillingAddress(initialAddress)
	returnAddress := pan.(*PAN).GetBillingAddress()

	if initialAddress != returnAddress {
		t.Log("Data: SetBillingAddress set the wrong value, or GetBillingAddress returned the wrong value")
		t.Fail()
	}
}
