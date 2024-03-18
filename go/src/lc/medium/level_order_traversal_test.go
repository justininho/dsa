package medium

import (
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 15}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 20}

	result := levelOrder(root)

	if len(result) != 3 {
		t.FailNow()
	}
}
