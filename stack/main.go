package main

import "fmt"

// Stack represents a basic stack data structure.
type Stack struct {
	items []interface{}
}

// Push Adds an element to the top of the stack.
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil 
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	// create a new stack
	myStack := Stack{}

	// Push elements
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Pop()

	// Pop and print elements
	for !myStack.IsEmpty() {
		fmt.Println(myStack.Pop())
	}
}
