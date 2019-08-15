package database

// Database is an interface that defines the methods within a database struct
//
type Database interface {
	SetIP(string)    // Set the IP of the Database struct
	GetIP() string   // Get the IP of the Database struct
	SetPort(string)  // Set the port of the Database struct
	GetPort() string // Get the port of the Database struct
}
