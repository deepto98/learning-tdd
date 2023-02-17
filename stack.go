package main

type Stack struct {
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Empty() bool {
	return true
}
