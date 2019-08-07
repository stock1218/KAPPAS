package model

// AmazonModel is a Model type struct that represents a model that uses AWS services
//
type AmazonModel struct{}

// NewAmazonModel is a function that returns a pointer to a new AmazonModel
//
func NewAmazonModel() *AmazonModel {
	return new(AmazonModel)
}
