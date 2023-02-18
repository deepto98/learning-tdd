package main

type Stack struct {
	size  int
	value int
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
	stack.value = val
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Pop() int {
	stack.size--
	return stack.value
}
