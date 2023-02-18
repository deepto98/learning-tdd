package main

import "testing"

func TestEmpty(t *testing.T) {

	stack := NewStack()

	// Assertions to check if stack is empty
	// Ideology : Empty should be written only after this case is defined
	if stack.Empty() != true {
		t.Error("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {

	stack := NewStack()

	// Add method is written after this case is defined
	stack.Add(1)

	if stack.Empty() != false {
		t.Error("Stack is empty")
	}
}
func TestSizeZero(t *testing.T) {
	stack := NewStack()

	if stack.Size() != 0 {
		t.Errorf("Expected : 0. ELements found : %d", stack.Size())
	}
}
