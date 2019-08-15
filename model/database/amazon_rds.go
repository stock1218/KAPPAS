package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
)

// AmazonRDS is a struct that represents a database that uses Amazon RDS.
//
type AmazonRDS struct {
	URL        string
	username   string
	password   string
	connection *sql.DB
}

// NewAmazonRDS is a function that returns a pointer to a new AmazonRDS struct.
//
func NewAmazonRDS(newURL string, newUsername string, newPassword string) *AmazonRDS {
	rds := new(AmazonRDS)

	rds.SetURL(newURL)
	rds.SetUsername(newUsername)
	rds.SetPassword(newPassword)

	newConn, err := sql.Open("postgres", rds.craftConnectionString())
	if err != nil {
		fmt.Println("Failed to connect to database: ", err)
	}

	rds.connection = newConn
	return rds
}

func (rds AmazonRDS) craftConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s?sslmode=disable", rds.GetUsername(), rds.GetPassword(), rds.GetURL())
}

// PutData will put Data into the database and return the data's ID.
//
func (rds AmazonRDS) PutData(newData string) string {
	var newID string
	err := rds.connection.QueryRow("INSERT INTO data (payload) VALUES ($1) RETURNING id", newData).Scan(&newID)
	if err != nil {
		fmt.Println("Error putting data: ", err)
	}
	return newID
}

// GetData will get the Data associated with the provided ID and return it from the database.
//
func (rds AmazonRDS) GetData(dataID string) string {
	var result string
	err := rds.connection.QueryRow("SELECT payload FROM data WHERE id=$1", dataID).Scan(&result)
	if err != nil {
		fmt.Println("Error getting data: ", err)
	}
	return result
}

// GetURL returns the IP of the database struct.
//
func (rds AmazonRDS) GetURL() string {
	return rds.URL
}

// SetURL sets the IP of the database struct.
//
func (rds *AmazonRDS) SetURL(newURL string) {
	rds.URL = newURL
}

// GetUsername will return the username currently being used to connect to the database.
//
func (rds AmazonRDS) GetUsername() string {
	return rds.username
}

// SetUsername will set the username used when connecting to the database.
//
func (rds *AmazonRDS) SetUsername(newUsername string) {
	rds.username = newUsername
}

// GetPassword will get the password used when connecting to the database.
//
func (rds AmazonRDS) GetPassword() string {
	return rds.password
}

// SetPassword will set the password used when connecting to the database.
//
func (rds *AmazonRDS) SetPassword(newPass string) {
	rds.password = newPass
}
