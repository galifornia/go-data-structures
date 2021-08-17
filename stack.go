package main

import "fmt"

type Stack struct {
	items []int
}

// push
func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

// pop
func (s *Stack) Pop() int {
	p := len(s.items) - 1
	item := s.items[len(s.items)-1]
	s.items = s.items[:p]
	return item
}

func main() {
	stack := &Stack{}
	stack.Push(44)
	stack.Push(24)
	stack.Push(6)
	stack.Push(36)

	fmt.Println(stack.items)

	stack.Pop()
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.items)
}
