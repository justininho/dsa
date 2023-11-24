package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()
	if s.head != nil {
		t.Errorf("NewStack = %v; want = %v", s.head, nil)
	}

	peek, ok := s.Peek()
	if ok || peek != 0 {
		t.Errorf("Peek = %v; want = %v", peek, 0)
		t.Errorf("Ok = %v; want = %v", ok, false)
	}

	s.Push(1)
	if s.head.value != 1 {
		t.Errorf("Push = %v; want = %v", s.head.value, 1)
	}

	peek, ok = s.Peek()
	if !ok || peek != 1 {
		t.Errorf("Peek = %v; want = %v", peek, 1)
		t.Errorf("Ok = %v; want = %v", ok, true)
	}

	length := s.Len()
	if length != 1 {
		t.Errorf("Len = %v; want = %v", length, 1)
	}

	s.Push(2)
	if s.head.value != 2 {
		t.Errorf("Push = %v; want = %v", s.head.value, 2)
	}

	peek, ok = s.Peek()
	if !ok || peek != 2 {
		t.Errorf("Peek = %v; want = %v", peek, 2)
		t.Errorf("Ok = %v; want = %v", ok, true)
	}

	length = s.Len()
	if length != 2 {
		t.Errorf("Len = %v; want = %v", length, 2)
	}

	value, ok := s.Pop()
	if !ok || value != 2 {
		t.Errorf("Pop = %v; want = %v", value, 2)
	}

	peek, ok = s.Peek()
	if !ok || peek != 1 {
		t.Errorf("Peek = %v; want = %v", peek, 1)
		t.Errorf("Ok = %v; want = %v", ok, true)
	}

	length = s.Len()
	if length != 1 {
		t.Errorf("Len = %v; want = %v", length, 1)
	}

	value, ok = s.Pop()
	if !ok || value != 1 {
		t.Errorf("Pop = %v; want = %v", value, 1)
	}

	peek, ok = s.Peek()
	if ok {
		t.Errorf("Peek = %v; want = %v", peek, 0)
		t.Errorf("Ok = %v; want = %v", ok, false)
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
