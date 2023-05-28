package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	if s.Head != nil {
		t.Errorf("NewStack = %v; want = %v", s.Head, nil)
	}

	s.Push(1)
	if s.Head.Value != 1 {
		t.Errorf("Push = %v; want = %v", s.Head.Value, 1)
	}

	s.Push(2)
	if s.Head.Value != 2 {
		t.Errorf("Push = %v; want = %v", s.Head.Value, 2)
	}

	popped, ok := s.Pop()
	if !ok || popped != 2 {
		t.Errorf("Pop = %v; want = %v", popped, 2)
	}

	popped, ok = s.Pop()
	if !ok || popped != 1 {
		t.Errorf("Pop = %v; want = %v", popped, 1)
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("Pop = %v; want = %v", ok, false)
	}
}
