package view

// RPC defines a struct for interacting with an RPC interface.
type RPC struct{}

// NewRPC creates a new RPC struct.
func NewRPC() *RPC {
	return new(RPC)
}
