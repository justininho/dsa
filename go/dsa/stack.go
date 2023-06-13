package dsa

import (
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Stack[T any] struct {
	Head *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Head: nil}
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{Value: value, Next: s.Head}
	s.Head = newNode
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Head == nil {
		var empty T
		return empty, false
	}
	value := s.Head.Value
	s.Head = s.Head.Next
	return value, true
}

func (s *Stack[t]) Print() {
	current := s.Head
	for current != nil {
		fmt.Printf("%v ", current.Value)
		current = current.Next
	}
	fmt.Println()
}
