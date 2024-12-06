package easy

import . "dsa/src/lc/lib"

/*
(111 | Easy)
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
Note: A leaf is a node with no children.
*/

// dfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minDepth(root.Left), minDepth(root.Right)

	if left == 0 || right == 0 {
		return max(left, right) + 1
	} else {
		return min(left, right) + 1
	}
}

// bfs
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	var queue = []*TreeNode{root}
//	depth := 1
//
//	for len(queue) > 0 {
//		n := len(queue)
//		for n > 0 {
//			n--
//			node := queue[0]
//			queue = queue[1:]
//
//			if node.Left == nil && node.Right == nil {
//				return depth
//			}
//
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//		depth++
//	}
//	return depth
//}
