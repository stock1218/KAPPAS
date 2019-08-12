package key

// Key is an interface that defines the methods required by Key structs
//
type Key interface {
	GetID() string // Return the ID of the key.
	SetID(string)  // Set the ID of the key.
}
