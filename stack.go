package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Size:", stack.Size()) // Size: 3

	item, _ := stack.Pop()
	fmt.Println("Popped item:", item) // Popped item: 3
	temp, err := stack.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Peek item:", temp) // Peek item: 2

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Is stack empty? false
}
