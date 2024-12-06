package easy

import . "dsa/src/lc/lib"

/*
(112 | Easy)
Given the root of a binary tree and an integer targetSum,
return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.
*/

// dfs
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

/*

type SumNode struct {
	sum  int
	node *TreeNode
}

// bfs
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var queue []*SumNode
	queue = append(queue, &SumNode{0, root})

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		c.sum += c.node.Val

		if c.node.Left == nil && c.node.Right == nil && c.sum == targetSum {
			return true
		}

		if c.node.Left != nil {
			queue = append(queue, &SumNode{c.sum, c.node.Left})
		}

		if c.node.Right != nil {
			queue = append(queue, &SumNode{c.sum, c.node.Right})
		}
	}
	return false
}

*/
