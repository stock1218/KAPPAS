package model

import "testing"

// TestNewModel tests if a Model variable can be created
func TestNewModel(t *testing.T) {
	var _ Model = nil
}

func TestNewAmazonModel(t *testing.T) {
	var server Model = NewAmazonModel()
	if server == nil {
		t.Log("Failed to create NewAmazonModel")
		t.Fail()
	}
}

// Declare the model object that will be used for the follwing tests
var server Model = NewAmazonModel()

// TestPutAndGetPAN tests if PutPAN will take a string, return an id, and GetPAN() will retrieve it
func TestPutAndGetPAN(t *testing.T) {
	data := "123"
	id, ok := server.PutPAN(data)

	if ok != nil {
		t.Log("Failed to call PutPAN: ", ok)
		t.Fail()
	}

	if id == "" {
		t.Log("PutPAN returned an empty string")
		t.Fail()
	}

	getData, ok := server.GetPAN(id)

	if ok != nil {
		t.Log("Error retrieving PAN: ", ok)
		t.Fail()
	}

	if data != getData {
		t.Log("Retrieved data was not correct")
		t.Fail()
	}
}

// TestEncrypt will test if encrypt() will take a plaintext string and encrypt it
func TestEncrypt(t *testing.T) {
	plaintext := "my secret"
	ciphertext, ok := server.encrypt(plaintext)

	if ok != nil {
		t.Log("Error encrypting data: ", ok)
		t.Fail()
	}

	if ciphertext == "" {
		t.Log("encrypt() returned an empty string")
	}
}
