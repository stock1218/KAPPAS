package main

import (
	"testing"
)

// TestStart will test if start() can be called.
func TestStart(t *testing.T) {
	ok := start()
	if ok != nil {
		t.Log("Failed to call start(): ", ok)
		t.Fail()
	}
}
