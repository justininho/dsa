package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traverse(root, &result)
	return result
}

func traverse(root *TreeNode, result *[]int) {
	if root != nil {
		traverse(root.Left, result)
		*result = append(*result, root.Val)
		traverse(root.Right, result)
	}
}
