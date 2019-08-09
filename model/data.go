package model

// Data defines an interface describing the methods required for a Data type.
//
type Data interface {
	ToJSON() []byte // Convert to a JSON byte array

	GetID() string // Get the ID of the pan
	SetID(string)  // Set the Id of the pan
}
