package easy

import (
	"dsa/src/lc/medium"
)

/*
(700 | Easy) searchBST
*/
func searchBST(root *medium.TreeNode, val int) *medium.TreeNode {
	if root == nil || root.Val == val {
		return root
	} else if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

//func bfs(node *TreeNode) []int {
//	var result []int
//	q := []*TreeNode{node}
//	for len(q) > 0 {
//		n := q[0]
//		q = q[1:]
//		if n != nil {
//			result = append(result, n.Val)
//			q = append(q, n.Left)
//			q = append(q, n.Right)
//		}
//	}
//	return result
//}
