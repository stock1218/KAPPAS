package database

import "testing"

// TestNewDatabase will test if a Database struct can be created.
//
func TestNewDatabase(t *testing.T) {
	var _ Database = nil
}

// TestNewAmazonRDS tests if an amazonRDS struct can be created.
//
func TestNewAmazonRDS(t *testing.T) {
	var database Database = NewAmazonRDS()

	if database == nil {
		t.Log("Database: failed to create AmazonRDS struct")
		t.Fail()
	}
}

// DatabaseSetUp will create a testable Database struct.
//
func DatabaseSetUp() Database {
	var db Database = NewAmazonRDS()
	return db
}

// TestSetAndGetIP tests if SetIP will set the ip of a database struct, and GetIP will return it.
//
func TestSetAndGetIP(t *testing.T) {
	db := DatabaseSetUp()

	initialIP := "127.0.0.1"

	db.SetIP(initialIP)

	returnedIP := db.GetIP()
	if initialIP != returnedIP {
		t.Log("Database: SetIP set the wrong value, or GetIP returned the wrong value")
		t.Fail()
	}
}

// TestSetAndGetPort will test if SetPort will set the port of a database struct, and GetPort will return it.
//
func TestSetAndGetPort(t *testing.T) {
	db := DatabaseSetUp()

	initialPort := "1337"

	db.SetPort(initialPort)

	returnedPort := db.GetPort()
	if initialPort != returnedPort {
		t.Log("Database: SetPort set the wrong value, or GetPort returned the wrong value")
		t.Fail()
	}
}
