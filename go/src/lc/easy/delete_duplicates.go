package easy

import . "dsa/src/lc/lib"

/*
83. Remove Duplicates from Sorted List
Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
Return the linked list sorted as well.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	var p *ListNode
	c := head
	for c != nil {
		if p != nil && p.Val == c.Val {
			c = c.Next
			p.Next = c
		} else {
			p = c
			c = c.Next
		}
	}
	return head
}
