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
		t.Errorf("Expected size : 0. ELements found : %d", stack.Size())
	}
}

func TestSizeOne(t *testing.T) {
	stack := NewStack()
	stack.Add(12)

	if stack.Size() != 1 {
		t.Error("Incorrect size")
		t.Log("Expected size : 1")
		t.Logf("Elements found : %d", stack.Size())
	}
}

func TestSizeThree(t *testing.T) {
	stack := NewStack()
	stack.Add(1)
	stack.Add(1)
	stack.Add(1)

	if stack.Size() != 3 {
		t.Error("Incorrect size returned")
		t.Log("Expected size : 3")
		t.Logf("Elements found : %d", stack.Size())
	}
}

func TestPopOne(t *testing.T) {
	stack := NewStack()
	stack.Add(11)
	val := stack.Pop()

	if val != 11 {
		t.Errorf("Expected 11. Received: %d", val)
	}

	if stack.Size() != 0 {
		t.Errorf("Expected size : 0. Received: %d", stack.Size())

	}
}

func TestPopTwo(t *testing.T) {
	stack := NewStack()
	stack.Add(11)
	stack.Add(20)
	val := stack.Pop()

	// LIFO - 20 has to be popped
	if val != 20 {
		t.Errorf("Expected 20. Received: %d", val)
	}

	if stack.Size() != 1 {
		t.Errorf("Expected size : 1. Received: %d", stack.Size())

	}
}
