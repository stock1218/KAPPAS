package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stock1218/KAPPAS/model/database"
	"github.com/stock1218/KAPPAS/model/key"
)

// AWSModel is a Model type struct that represents a model that uses AWS services.
//
type AWSModel struct {
	database database.Database
	key      key.Key
}

var pan string

// NewAWSModel is a function that returns a pointer to a new AWSModel.
//
func NewAWSModel() *AWSModel {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1")},
	)
	model := new(AWSModel)
	model.SetDatabase(database.NewAmazonRDS())
	model.SetKey(key.NewKMS(*sess))
	return model
}

// PutPAN takes the data to be saved, and returns the id of the saved data, and an error if there is one.
//
func (model AWSModel) PutPAN(data string) (string, error) {
	pan = data
	return "1", nil
}

// GetPAN takes the ID of the PAN being fetched, and returns the saved data, and an error if there is one.
//
func (model AWSModel) GetPAN(id string) (string, error) {
	return pan, nil
}

// encrypt takes a string of plaintext to encrypt, and returns encrypted ciphertext.
//
func (model AWSModel) encrypt(data string) (string, error) {
	return "ciphertext", nil
}

// decrypt takes a string of ciphertext to decrypt, and returns the corrisponding plaintext.
//
func (model AWSModel) decrypt(ciphertext string) (string, error) {
	return "my secret", nil
}

// GetDatabase gets a copy of the database currently being used.
//
func (model AWSModel) GetDatabase() database.Database {
	return model.database
}

// SetDatabase sets the pointer to the database currently being used.
//
func (model *AWSModel) SetDatabase(newDB database.Database) {
	model.database = newDB
}

// GetKey will return a copy of the key currently being used.
//
func (model AWSModel) GetKey() key.Key {
	return model.key
}

// SetKey sets the pointer to the database currently being used.
//
func (model *AWSModel) SetKey(newKey key.Key) {
	model.key = newKey
}
