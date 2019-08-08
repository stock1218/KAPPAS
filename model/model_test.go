package model

import (
	"testing"
)

// TestNewModel tests if a Model variable can be created.
func TestNewModel(t *testing.T) {
	var _ Model = nil
}

func TestNewAWSModel(t *testing.T) {
	var server Model = NewAWSModel()
	if server == nil {
		t.Log("Failed to create NewAWSModel")
		t.Fail()
	}
}

// Declare the model object that will be used for the follwing tests.
var server Model = NewAWSModel()

// TestPutAndGetPAN tests if PutPAN will take a string, return an id, and GetPAN() will retrieve it.
func TestPutAndGetPAN(t *testing.T) {
	data := "123"
	id, ok := server.PutPAN(data)

	if ok != nil {
		t.Log("Failed to call PutPAN: ", ok)
		t.Fail()
	} else if id == "" {
		t.Log("PutPAN returned an empty string")
		t.Fail()
	}

	getData, ok := server.GetPAN(id)

	if ok != nil {
		t.Log("Error retrieving PAN: ", ok)
		t.Fail()
	} else if data != getData {
		t.Log("Retrieved data was not correct")
		t.Fail()
	}
}

// TestEncryptAndDecrypt will test if encrypt() will take a plaintext string and encrypt it,
// and decrypt() will decrypt the ciphertext back into plaintext.
func TestEncryptAndDecrypt(t *testing.T) {
	plaintext := "my secret"
	ciphertext, ok := server.encrypt(plaintext)

	if ok != nil {
		t.Log("Error encrypting data: ", ok)
		t.Fail()
	} else if ciphertext == "" {
		t.Log("encrypt() returned an empty string")
	}

	decrypted, ok := server.decrypt(ciphertext)

	if ok != nil {
		t.Log("Error decrypting data: ", ok)
		t.Fail()
	} else if decrypted != plaintext {
		t.Log("Decrypt() returned wrong plaintext")
		t.Fail()
	}

}

// TestSetAndGetDatabase will test to see if SetDatabase will set the database the AWSModel is uing, and GetDatabse will return it.
func TestSetAndGetDatabase(t *testing.T) {
	testDB := NewAmazonRDS()
	server.(*AWSModel).SetDatabase(testDB)
	database := server.(*AWSModel).GetDatabase() // Cast server to a pointer to an AWSModel and call GetDatabase

	if database == nil {
		t.Log("GetDatabase() returned nil")
		t.Fail()
	} else if testDB != database {
		t.Log("GetDatabase returned incorrect value, or SetDatabase set incorrect value")
		t.Fail()
	}
}

// TestSetAndGetKey will test if SetKey will set the key used by the AWSModel, and GetKey will return it back.
func TestSetAndGetKey(t *testing.T) {
	myKey := NewKMS()
	server.(*AWSModel).SetKey(myKey)
	getKey := server.(*AWSModel).GetKey()
	if getKey == nil {
		t.Log("GetKey returned nil")
		t.Fail()
	} else if myKey != getKey {
		t.Log("GetKey returned incorrect value, or SetKey set incorrect value")
	}
}
