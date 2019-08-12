package data

import (
	"testing"
	"time"
)

// TestNewData will test if a data variable can be created.
//
func TestNewData(t *testing.T) {
	var _ Data = nil
}

// TestNewPANModel tests if a PAN type Data variable can be created.
//
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

// TestSentAndGetBillingAddress tests if SetBillingAddress will set the billing address of a pan,
// and GetBillingAddress will return it.
//
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

// TestToJSON tests if a pan struct can be converted to and from JSON.
//
func TestToJSON(t *testing.T) {
	pan := PANSetUp()
	pan.SetID("123")

	returned := pan.ToJSON()

	expected := `{"id":"123","cardNumber":"","cardHolder":"","experationDate":"0001-01-01T00:00:00Z","billingAddress":""}`
	if string(returned) != expected {
		t.Log("Data: ToJSON failed to return the correct JSON")
		t.Fail()
	}
}

// TestFromJSON tests if a pan struct can be created from a JSON string.
//
func TestFromJSON(t *testing.T) {
	expectedPan := PANSetUp()
	expectedPan.SetID("123")

	panString := `{"id":"123","cardNumber":"","cardHolder":"","experationDate":"0001-01-01T00:00:00Z","billingAddress":""}`

	testPan := NewPANFromJSON(panString)

	if expectedPan.GetID() != testPan.GetID() {
		t.Log("Data: FromJSON didn't set the correct values")
		t.Fail()
	}
}
