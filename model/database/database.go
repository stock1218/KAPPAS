package database

// Database is an interface that defines the methods within a database struct
//
type Database interface {
	PutData(string) string // Save data to Database and return it's ID
	GetData(string) string // Get data associated with the ID from the Database
	SetURL(string)         // Set the URL of the Database struct
	GetURL() string        // Get the URL of the Database struct
	SetUsername(string)    // Set the username to use when connecting to the Database
	GetUsername() string   // Get the username used by the Database struct
	SetPassword(string)    // Set the password to use when connecting to the Database
	GetPassword() string   // Get the password used by the Database struct
}
