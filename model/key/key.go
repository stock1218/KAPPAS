package key

// Key is an interface that defines the methods required by Key structs
//
type Key interface {
	Encrypt(string) string // Encrypt a string
	Decrypt(string) string // Decrypt a string
	GetID() string         // Return the ID of the key.
	SetID(string)          // Set the ID of the key.
}
