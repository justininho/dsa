package easy

import (
	. "dsa/src/lc/lib"
	"testing"
)

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		input    *ListNode
		expected bool
	}{
		{&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}, true},
		{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, true},
		{&ListNode{Val: 1}, false},
	}

	for _, tt := range tests {
		actual := hasCycle(tt.input)
		if actual != tt.expected {
			t.Errorf("hasCycle(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}
