package easy

import (
	. "dsa/src/lc/lib"
	"testing"
)

func TestMergeList(t *testing.T) {
	var tests = []struct {
		input1   *ListNode
		input2   *ListNode
		expected *ListNode
	}{
		//{&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}},
		{nil, &ListNode{Val: 0}, &ListNode{Val: 0}},
	}

	for _, tt := range tests {
		actual := mergeTwoLists(tt.input1, tt.input2)
		if !IsSameList(actual, tt.expected) {
			t.Errorf("mergeTwoLists(%v, %v): expected %v, actual %v", tt.input1, tt.input2, tt.expected, actual)
		}
	}
}
