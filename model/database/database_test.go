package database

import (
	"testing"
)

// TestNewDatabase will test if a Database struct can be created.
//
func TestNewDatabase(t *testing.T) {
	var _ Database = nil
}

// TestNewAmazonRDS tests if an amazonRDS struct can be created.
//
func TestNewAmazonRDS(t *testing.T) {
	var database Database = NewAmazonRDS("localhost/kappas", "stock1218", "")

	if database == nil {
		t.Log("Database: failed to create AmazonRDS struct")
		t.Fail()
	}
}

// DatabaseSetUp will create a testable Database struct.
//
func DatabaseSetUp() Database {
	var db Database = NewAmazonRDS("localhost/kappas", "stock1218", "")
	return db
}

// TestSetAndGetIP tests if SetIP will set the ip of a database struct, and GetIP will return it.
//
func TestSetAndGetIP(t *testing.T) {
	db := DatabaseSetUp()

	initialURL := "127.0.0.1"

	db.SetURL(initialURL)

	returnedURL := db.GetURL()
	if initialURL != returnedURL {
		t.Log("Database: SetIP set the wrong value, or GetIP returned the wrong value")
		t.Fail()
	}
}

// TestSetAndGetUsername tests if SetUsername will set the username of a database struct, and GetUsername will return it.
//
func TestSetAndGetUsername(t *testing.T) {
	db := DatabaseSetUp()

	initialUsername := "Clu"

	db.SetUsername(initialUsername)

	returnedUsername := db.GetUsername()
	if initialUsername != returnedUsername {
		t.Log("Database: SetUsername set the wrong value, or GetUsername returned the wrong value")
		t.Fail()
	}
}

// TestSetAndGetPassword tests if SetPassword will set the password of a database struct, and GetPassword will return it.
//
func TestSetAndGetPassword(t *testing.T) {
	db := DatabaseSetUp()

	initialPassword := "I fight for the users"

	db.SetPassword(initialPassword)

	returnedPassword := db.GetPassword()
	if initialPassword != returnedPassword {
		t.Log("Database: SetPassword set the wrong value, or GetPassword returned the wrong value")
		t.Fail()
	}
}

// TestPutAndGetData tests if PutData will save data to the database, and GetData will return it.
//
func TestPutAngGetData(t *testing.T) {
	db := DatabaseSetUp()

	initialData := "Biodigital Jazz"
	id := db.PutData(initialData)

	returnedData := db.GetData(id)
	if initialData != returnedData {
		t.Log("Database: PutData didn't put the correct value into the database, or GetData didn't return the correct value")
		t.Fail()
	}

}
