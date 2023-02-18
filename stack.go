package main

type Stack struct {
	isEmpty bool
	size    int
}

func NewStack() *Stack {
	return &Stack{
		isEmpty: true,
		size:    0,
	}
}

func (stack *Stack) Empty() bool {
	return stack.isEmpty
}

func (stack *Stack) Add(val int) {
	stack.isEmpty = false
	stack.size++
}

func (stack *Stack) Size() int {
	return stack.size
}
