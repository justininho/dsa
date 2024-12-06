package easy

import (
	. "dsa/src/lc/lib"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
	}

	for _, tt := range tests {
		if got := deleteDuplicates(tt.head); !IsSameList(got, tt.expected) {
			t.Errorf("deleteDuplicates(%v), got=%v, expected=%v", tt.head, got, tt.expected)
		}
	}
}
