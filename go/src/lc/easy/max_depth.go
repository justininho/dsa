package easy

import (
	. "dsa/src/lc/lib"
)

/*
(104 | Easy) Maximum Depth of Binary Tree
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
