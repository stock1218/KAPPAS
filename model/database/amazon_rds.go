package database

// AmazonRDS is a struct that represents a database that uses Amazon RDS.
//
type AmazonRDS struct {
	IP       string
	port     string
	username string
	password string
}

// NewAmazonRDS is a function that returns a pointer to a new AmazonRDS struct.
//
func NewAmazonRDS() *AmazonRDS {
	return new(AmazonRDS)
}

// GetIP returns the IP of the database struct.
//
func (rds AmazonRDS) GetIP() string {
	return rds.IP
}

// SetIP sets the IP of the database struct.
//
func (rds *AmazonRDS) SetIP(newIP string) {
	rds.IP = newIP
}

// GetPort returns the port of the database struct.
//
func (rds AmazonRDS) GetPort() string {
	return rds.port
}

// SetPort will set the port used to connect to the underlying database.
//
func (rds *AmazonRDS) SetPort(newPort string) {
	rds.port = newPort
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
