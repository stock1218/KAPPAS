package model

// AmazonModel is a Model type struct that represents a model that uses AWS services
//
type AmazonModel struct{}

var pan string

// NewAmazonModel is a function that returns a pointer to a new AmazonModel
//
func NewAmazonModel() *AmazonModel {
	return new(AmazonModel)
}

// PutPAN takes the data to be saved, and returns the id of the saved data, and an error if there is one
//
func (model AmazonModel) PutPAN(data string) (string, error) {
	pan = data
	return "1", nil
}

// GetPAN takes the ID of the PAN being fetched, and returns the saved data, and an error if there is one
//
func (model AmazonModel) GetPAN(id string) (string, error) {
	return pan, nil
}
