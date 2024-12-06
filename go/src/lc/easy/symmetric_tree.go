package easy

import . "dsa/src/lc/lib"

/*
(101 | Easy)
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return nodeEqual(root.Left, root.Right)
}

func nodeEqual(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return nodeEqual(p.Left, q.Right) && nodeEqual(p.Right, q.Left)
}
