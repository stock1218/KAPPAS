package key

// KMS represents an abstraction of the AWS key management system (KMS).
//
type KMS struct {
	ID string
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
