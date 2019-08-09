package model

import (
	"testing"
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

// TestSetAndIsEncrypted tests if SetEncrypted will set a Data struct's encrypted state, and
// IsEncrypted will return it.
func TestSetAndIsEncrypted(t *testing.T) {
	pan := PANSetUp()
	myState := true
	pan.SetEncrypted(myState)
	returnedState := pan.IsEncrypted()

	if myState != returnedState {
		t.Log("Data: SetEncrypted set the wrong state or IsEncrypted returned the wrong state")
		t.Fail()
	}

}
