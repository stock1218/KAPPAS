package view

// RPC defines a struct for interacting with an RPC interface.
type RPC struct {
	observers []func(string)
}

// NewRPC creates a new RPC struct.
//
func NewRPC() *RPC {
	return new(RPC)
}

// GetObservers will return the array of observer functions registered to the RPC.
//
func (rpc RPC) GetObservers() []func(string) {
	return rpc.observers
}
