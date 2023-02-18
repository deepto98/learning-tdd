package main

type Stack struct {
	isEmpty bool
}

func NewStack() *Stack {
	return &Stack{
		isEmpty: true,
	}
}

func (stack *Stack) Empty() bool {
	return stack.isEmpty
}

func (stack *Stack) Add(val int) {
	stack.isEmpty = false
}

func (stack *Stack) Size() int {

	if !stack.isEmpty {
		return 1
	}

	return 0
}
