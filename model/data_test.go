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
func PANSetUp() *PAN {
	return NewPAN()
}

// TestGetPayload tests if GetPayload will return a string.
//
func TestGetPayload(t *testing.T) {
	pan := PANSetUp()
	payload := pan.GetPayload()

	if payload == "" {
		t.Log("Data: GetPayload returned empty string")
		t.Fail()
	}
}
