package model

import "testing"

// TestNewKey tests if a Key type variable can be created.
//
func TestNewKey(t *testing.T) {
	var _ Key = nil
}

// TestNewKMS tests if a KMS struct can be created.
//
func TestNewKMS(t *testing.T) {
	var key Key = NewKMS()
	if key == nil {
		t.Log("Key: failed to create new KMS")
		t.Fail()
	}
}
