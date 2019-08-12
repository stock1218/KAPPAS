package key

import (
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
func (kms KMS) GetID() string {
	return kms.ID
}

// SetID sets the ID of the KMS struct.
//
func (kms *KMS) SetID(newID string) {
	kms.ID = newID
}

// GetClient gets the Session of the KMS struct.
//
func (kms KMS) GetClient() kms.KMS {
	return kms.client
}

// SetClient gets the Session of the KMS struct.
//
func (kms *KMS) SetClient(newClient kms.KMS) {
	kms.client = newClient
}
