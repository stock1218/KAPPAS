package key

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMS represents an abstraction of the AWS key management system (KMS).
//
type KMS struct {
	ID     string
	client kms.KMS // KMS struct from aws-sdk-go
}

// NewKMS returns a pointer to a new KMS struct.
//
func NewKMS(sess session.Session) *KMS {
	newKMS := new(KMS)
	newClient := kms.New(&sess)
	newKMS.client = *newClient
	return newKMS
}

// GetID gets the ID of the KMS struct.
//
func (key KMS) GetID() string {
	return key.ID
}

// SetID sets the ID of the KMS struct.
//
func (key *KMS) SetID(newID string) {
	key.ID = newID
}

// GetClient gets the Session of the KMS struct.
//
func (key KMS) GetClient() kms.KMS {
	return key.client
}

// SetClient gets the Session of the KMS struct.
//
func (key *KMS) SetClient(newClient kms.KMS) {
	key.client = newClient
}

// Encrypt encrypts a string using the AWS KMS service.
//
func (key KMS) Encrypt(plaintext string) string {
	result, _ := key.client.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(key.ID),
		Plaintext: []byte(plaintext),
	})

	return result.GoString()
}

// Decrypt decrypts a string using the AWS KMS service.
//
func (key KMS) Decrypt(ciphertext string) string {
	result, err := key.client.Decrypt(&kms.DecryptInput{
		CiphertextBlob: []byte(ciphertext),
	})

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return result.GoString()
}
