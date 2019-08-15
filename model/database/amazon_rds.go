package database

// AmazonRDS is a struct that represents a database that uses Amazon RDS.
//
type AmazonRDS struct {
	URL      string
	username string
	password string
}

// NewAmazonRDS is a function that returns a pointer to a new AmazonRDS struct.
//
func NewAmazonRDS() *AmazonRDS {
	return new(AmazonRDS)
}

// PutData will put Data into the database and return the data's ID.
//
func (rds AmazonRDS) PutData(newData string) string {
	return ""
}

// GetData will get the Data associated with the provided ID and return it from the database.
//
func (rds AmazonRDS) GetData(dataID string) string {
	return "Biodigital Jazz"
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
