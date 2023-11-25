package lc

import "testing"

func TestReverseList(t *testing.T) {
	var tests = []struct {
		input    *ListNode
		expected *ListNode
	}{
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		//{&ListNode{1, &ListNode{2, nil}}, &ListNode{2, &ListNode{1, nil}}},
		//{&ListNode{1, nil}, &ListNode{1, nil}},
		//{nil, nil},
	}

	for _, tt := range tests {
		actual := reverseList(tt.input)
		if !isSameList(actual, tt.expected) {
			t.Errorf("reverseList(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}

}

func isSameList(actual *ListNode, expected *ListNode) bool {
	if actual == nil && expected == nil {
		return true
	}

	if actual == nil || expected == nil {
		return false
	}

	if actual.Val != expected.Val {
		return false
	}

	return isSameList(actual.Next, expected.Next)
}
