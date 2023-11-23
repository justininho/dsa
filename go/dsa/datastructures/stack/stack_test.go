package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()
	if s.Head != nil {
		t.Errorf("NewStack = %v; want = %v", s.Head, nil)
	}

	s.Push(1)
	if s.Head.Value != 1 {
		t.Errorf("Push = %v; want = %v", s.Head.Value, 1)
	}

	length := s.Len()
	if length != 1 {
		t.Errorf("Len = %v; want = %v", length, 1)
	}

	s.Push(2)
	if s.Head.Value != 2 {
		t.Errorf("Push = %v; want = %v", s.Head.Value, 2)
	}

	length = s.Len()
	if length != 2 {
		t.Errorf("Len = %v; want = %v", length, 2)
	}

	value, ok := s.Pop()
	if !ok || value != 2 {
		t.Errorf("Pop = %v; want = %v", value, 2)
	}

	length = s.Len()
	if length != 1 {
		t.Errorf("Len = %v; want = %v", length, 1)
	}

	value, ok = s.Pop()
	if !ok || value != 1 {
		t.Errorf("Pop = %v; want = %v", value, 1)
	}

	length = s.Len()
	if length != 0 {
		t.Errorf("Len = %v; want = %v", length, 0)
	}

	_, ok = s.Pop()
	if ok {
		t.Errorf("Pop = %v; want = %v", ok, false)
	}
}
