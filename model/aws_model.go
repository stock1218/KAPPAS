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

// NewAWSModel is a function that returns a pointer to a new AWSModel.
//
func NewAWSModel() *AWSModel {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	model := new(AWSModel)
	model.SetDatabase(database.NewAmazonRDS("localhost/kappas", "stock1218", ""))
	newKey := key.NewKMS(*sess)
	newKey.SetID("arn:aws:kms:us-east-1:188809012751:key/e33ea72b-0ffd-4ba3-aa12-bebced4bd23d")
	model.SetKey(newKey)
	return model
}

// PutPAN takes the data to be saved, and returns the id of the saved data, and an error if there is one.
//
func (model AWSModel) PutPAN(data string) (string, error) {
	ciphertext := model.key.Encrypt(data)
	newID := model.database.PutData(ciphertext)
	return newID, nil
}

// GetPAN takes the ID of the PAN being fetched, and returns the saved data, and an error if there is one.
//
func (model AWSModel) GetPAN(id string) (string, error) {
	ciphertext := model.database.GetData(id)
	plaintext := model.key.Decrypt(ciphertext)
	return plaintext, nil

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
