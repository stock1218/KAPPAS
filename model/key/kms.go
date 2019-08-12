package key

// KMS represents an abstraction of the AWS key management system (KMS).
//
type KMS struct{}

// NewKMS returns a pointer to a new KMS struct.
//
func NewKMS() *KMS {
	return new(KMS)
}
