package easy

import (
	. "dsa/src/lc/lib"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 15}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 20}

	result := maxDepth(root)

	if result != 3 {
		t.FailNow()
	}
}
