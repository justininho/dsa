package easy

import . "dsa/src/lc/lib"

/**
 * Definition for singly-linked list.
 */

// (206 | Easy) reverseList
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
