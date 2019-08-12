package key

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

// KeySetUp is used as a setup function for creating a testable KMS.
//
func KeySetUp() Key {
	var key Key = NewKMS()
	return key
}

// TestSetAndGetID tests if SetID will set the ID of a key, and GetID will return it.
//
func TestSetAndGetID(t *testing.T) {
	key := KeySetUp()

	initialID := "123"
	key.SetID("123")
	returnedID := key.GetID()

	if initialID != returnedID {
		t.Log("SetID set the wrong value, or GetID returned the wrong value.")
		t.Fail()
	}

}
