package model

// AmazonRDS is a struct that represents a database that uses Amazon RDS.
//
type AmazonRDS struct{}

// NewAmazonRDS is a function that returns a pointer to a new AmazonRDS struct.
//
func NewAmazonRDS() *AmazonRDS {
	return new(AmazonRDS)
}
