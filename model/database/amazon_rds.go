package database

// AmazonRDS is a struct that represents a database that uses Amazon RDS.
//
type AmazonRDS struct {
	IP string
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
