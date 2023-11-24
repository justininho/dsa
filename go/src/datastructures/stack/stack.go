package stack

import (
	"fmt"
)

// node is a node in a Stack
type node[T any] struct {
	value T
	next  *node[T]
}

// Stack is a LIFO data structure
type Stack[T any] struct {
	head *node[T]
}

// New Creates a new Stack
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds a new value to the beginning of the Stack
func (s *Stack[T]) Push(value T) {
	newNode := &node[T]{value, s.head}
	s.head = newNode
}

// Pop removes the value on top of the Stack and returns it
func (s *Stack[T]) Pop() (T, bool) {
	if s.head == nil {
		var empty T
		return empty, false
	}
	value := s.head.value
	s.head = s.head.next
	return value, true
}

// Peek returns the value of on top of the Stack without popping
func (s *Stack[T]) Peek() (T, bool) {
	if head := s.head; head != nil {
		return s.head.value, true
	}
	var empty T
	return empty, false
}

// Len returns the length of the Stack
func (s *Stack[T]) Len() int {
	length, node := 0, s.head
	for node != nil {
		length++
		node = node.next
	}
	return length
}

// Print prints the Stack
func (s *Stack[t]) Print() {
	current := s.head
	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}
	fmt.Println()
}
