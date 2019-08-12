package key

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

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

// TestSetAndGetSession tests if SetSession will set the session for a KMS struct, and GetSession will return it.
//
func TestSetAndGetSession(t *testing.T) {
	key := KeySetUp()

	initialSess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)

	key.(*KMS).SetSession(*initialSess)

	returnedSess := key.(*KMS).GetSession()
	if initialSess.Config.Region != returnedSess.Config.Region {
		t.Log("SetSession set the incorrect value, or GetSession returned the wrong value")
		t.Fail()
	}
}
