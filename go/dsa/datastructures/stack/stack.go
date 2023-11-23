package stack

import (
	"fmt"
)

// node is a node in a Stack
type node[T any] struct {
	Value T
	Next  *node[T]
}

// Stack is a LIFO data structure
type Stack[T any] struct {
	Head *node[T]
}

// New Creates a new Stack
func New[T any]() *Stack[T] {
	return &Stack[T]{Head: nil}
}

// Peek returns the value of on top of the Stack without popping
func (s Stack[T]) Peek() (T, bool) {
	if head := s.Head; head != nil {
		return s.Head.Value, true
	}
	var empty T
	return empty, false
}

// Push adds a new value to the beginning of the Stack
func (s *Stack[T]) Push(value T) {
	newNode := &node[T]{Value: value, Next: s.Head}
	s.Head = newNode
}

// Pop removes the value on top of the Stack and returns it
func (s *Stack[T]) Pop() (T, bool) {
	if s.Head == nil {
		var empty T
		return empty, false
	}
	value := s.Head.Value
	s.Head = s.Head.Next
	return value, true
}

// Len returns the length of the Stack
func (s Stack[T]) Len() int {
	length, node := 0, s.Head
	for node != nil {
		length++
		node = node.Next
	}
	return length
}

// Print prints the Stack
func (s *Stack[t]) Print() {
	current := s.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
