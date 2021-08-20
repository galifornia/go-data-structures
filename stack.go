package main

import (
	"errors"
)

type Stack struct {
	items []interface{} // FIXME: use generics
}

// push
func (s *Stack) Push(v interface{}) {
	s.items = append(s.items, v)
}

// pop
func (s *Stack) Pop() (interface{}, error) {
	p := len(s.items) - 1
	if p < 0 {
		return -1, errors.New("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:p]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) <= 0
}

// func main() {
// 	stack := &Stack{}
// 	stack.Push(44)
// 	stack.Push(24)
// 	stack.Push(6)
// 	stack.Push(36)

// 	fmt.Println(stack.items)

// 	stack.Pop()
// 	stack.Pop()
// 	stack.Pop()
// 	stack.Pop()
// 	stack.Pop()
// 	stack.Pop()
// 	fmt.Println(stack.items)
// }
