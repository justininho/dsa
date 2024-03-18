package lib

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsSameList(actual *ListNode, expected *ListNode) bool {
	if actual == nil && expected == nil {
		return true
	}

	if actual == nil || expected == nil {
		return false
	}

	if actual.Val != expected.Val {
		return false
	}

	return IsSameList(actual.Next, expected.Next)
}
