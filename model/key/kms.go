package key

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

// KMS represents an abstraction of the AWS key management system (KMS).
//
type KMS struct {
	ID   string
	sess session.Session
}

// NewKMS returns a pointer to a new KMS struct.
//
func NewKMS() *KMS {
	return new(KMS)
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

// GetSession gets the Session of the KMS struct.
//
func (kms KMS) GetSession() session.Session {
	return kms.sess
}

// SetSession gets the Session of the KMS struct.
//
func (kms *KMS) SetSession(newSess session.Session) {
	kms.sess = newSess
}
