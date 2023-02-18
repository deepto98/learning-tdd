package main

type Stack struct {
	size   int
	values []int
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
	stack.values = append(stack.values, val)
	stack.size++
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Pop() int {
	stack.size--
	return stack.values[stack.size]
}
