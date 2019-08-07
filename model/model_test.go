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

// TestGetPAN tests if PutPAN will take a string, and save it to a database
func TestGetPAN(t *testing.T) {
	ok := server.PutPAN("123")

	if ok != nil {
		t.Log("Failed to call PutPAN, ", ok)
		t.Fail()
	}
}
