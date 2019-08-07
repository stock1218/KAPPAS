package model

// Model is an interface that defines the functions for Model types
//
type Model interface {
	PutPAN(string) (string, error)  // Put a PAN value into a model
	GetPAN(string) (string, error)  // Get a PAN value in a model
	encrypt(string) (string, error) // Encrypt a string
	decrypt(string) (string, error) // Decrypt a string encrypted by this model
}
