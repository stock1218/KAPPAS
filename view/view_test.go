package view

import "testing"

// Test is a View type can be created.
func TestNewview(t *testing.T) {
	var _ View = nil
}

// Test if an RPC type View can be created.
func TestNewRPC(t *testing.T) {
	var rpc View = NewRPC()

	if rpc == nil {
		t.Log("View: Failed to create RPC view")
		t.Fail()
	}
}

// A setup function for the following tests to create an RPC struct.
func viewSetup() *RPC {
	return NewRPC()
}

// Test if GetObservers will register an observer to a view type.
func TestGetObservers(t *testing.T) {
	testRPC := viewSetup()

	_ = testRPC.GetObservers()
}
