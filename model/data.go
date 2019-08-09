package model

// Data defines an interface describing the methods required for a Data type.
//
type Data interface {
	GetID() string // Get the ID of the pan
	SetID(string)  // Set the Id of the pan
}
