package easy

import (
	"dsa/src/lc/medium"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &medium.TreeNode{Val: 3}
	root.Left = &medium.TreeNode{Val: 9}
	root.Left.Left = &medium.TreeNode{Val: 15}
	root.Left.Right = &medium.TreeNode{Val: 7}
	root.Right = &medium.TreeNode{Val: 20}

	result := maxDepth(root)

	if result != 3 {
		t.FailNow()
	}
}
