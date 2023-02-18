package main

type Stack struct {
	size int
}

func NewStack() *Stack {
	return &Stack{
		size: 0,
	}
}

func (stack *Stack) Empty() bool {
	return stack.size == 0
}

func (stack *Stack) Add(val int) {
	stack.size++
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Pop() int {
	return 100
}
