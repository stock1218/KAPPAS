package database

import "testing"

// TestNewDatabase will test if a Database struct can be created.
//
func TestNewDatabase(t *testing.T) {
	var _ Database = nil
}

// TestNewAmazonRDS tests if an amazonRDS struct can be created
//
func TestNewAmazonRDS(t *testing.T) {
	var database Database = NewAmazonRDS()

	if database == nil {
		t.Log("Database: failed to create AmazonRDS struct")
		t.Fail()
	}
}
