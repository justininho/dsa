package medium

import "testing"

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	if ms.GetMin() != -3 {
		t.Errorf("GetMin = %v; want = %v", ms.GetMin(), -3)
	}
	ms.Pop()
	if ms.Top() != 0 {
		t.Errorf("Top = %v; want = %v", ms.Top(), 0)
	}
	if ms.GetMin() != -2 {
		t.Errorf("GetMin = %v; want = %v", ms.GetMin(), -2)
	}
}
