package view

import "testing"

// Test is a View type can be created.
func TestNewView(t *testing.T) {
	var _ View = nil
}

// Test is an RPC type View can be created.
func TestNewRPC(t *testing.T) {
	var rpc View = NewRPC()

	if rpc == nil {
		t.Log("View: Failed to create RPC view")
		t.Fail()
	}
}
