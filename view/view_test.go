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

// Test if GetObservers will return the observers registered to a view type.
func TestGetObservers(t *testing.T) {
	testRPC := viewSetup()

	_ = testRPC.GetObservers()
}

// Test if Register will register an observer function to a view type.
func TestRegister(t *testing.T) {
	testRPC := viewSetup()

	f := func(data string) {}
	testRPC.Register(f)

	observers := testRPC.GetObservers()

	if observers == nil {
		t.Log("View: rpc returned nil")
		t.Fail()
	}

}
