package model

// Data defines an interface describing the methods required for a Data type.
//
type Data interface {
	GetPayload() string // Get the payload contained in Data
	SetPayload(string)  // Set the payload contained in Data
	GetID() string      // Get the ID of the pan
	SetID(string)       // Set the Id of the pan
}
