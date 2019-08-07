package model

import "testing"

// Test if a Model variable can be created
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
