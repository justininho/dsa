package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int]()

	peek, ok := q.Peek()
	if ok {
		t.Errorf("peek = %v, want = %v", peek, 0)
		t.Errorf("ok = %v, want = %v", ok, false)
	}

	q.Enqueue(1)
	length := q.Len()
	if length != 1 {
		t.Errorf("length = %v, want = %v", length, 1)
	}

	q.Enqueue(2)
	length = q.Len()
	if length != 2 {
		t.Errorf("length = %v, want = %v", length, 2)
	}

	q.Enqueue(3)
	length = q.Len()
	if length != 3 {
		t.Errorf("length = %v, want = %v", length, 3)
	}

	rm, ok := q.Dequeue()
	if !ok || rm != 1 {
		t.Errorf("rm = %v, want = %v", rm, 1)
		t.Errorf("ok = %v, want = %v", ok, true)
	}

	length = q.Len()
	if length != 2 {
		t.Errorf("length = %v, want = %v", length, 2)
	}

	rm, ok = q.Dequeue()
	if !ok || rm != 2 {
		t.Errorf("rm = %v, want = %v", rm, 2)
		t.Errorf("ok = %v, want = %v", ok, true)
	}

	length = q.Len()
	if length != 1 {
		t.Errorf("length = %v, want = %v", length, 1)
	}

	rm, ok = q.Dequeue()
	if !ok || rm != 3 {
		t.Errorf("rm = %v, want = %v", rm, 3)
		t.Errorf("ok = %v, want = %v", ok, true)
	}

	length = q.Len()
	if length != 0 {
		t.Errorf("length = %v, want = %v", length, 0)
	}

	rm, ok = q.Dequeue()
	if ok {
		t.Errorf("ok = %v, want = %v", ok, false)
	}
}
