package easy

import (
	"dsa/src/lc/medium"
)

/*
(104 | Easy) Maximum Depth of Binary Tree
*/
func maxDepth(root *medium.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
