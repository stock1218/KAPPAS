package model

// Model is an interface that defines the functions for Model types
//
type Model interface {
	PutPAN(string) error
}
