package easy

import . "dsa/src/lc/lib"

/*
(100 | Easy)
Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil || q == nil && p != nil {
		return false
	} else if p == nil && q == nil {
		return true
	}

	if p == nil && q == nil || p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
