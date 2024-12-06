package easy

import (
	. "dsa/src/lc/lib"
	"math"
)

/*
(110 | Easy) Given a binary tree, determine if it is height-balanced.
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)
	diff := depthLeft - depthRight
	if math.Abs(float64(diff)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}
