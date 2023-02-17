package main

import "testing"

func TestEmpty(t *testing.T) {

	stack := NewStack()

	// Assertions to check if stack is empty
	if stack.Empty() != true {
		t.Error("Stack was not empty")
	}
}
