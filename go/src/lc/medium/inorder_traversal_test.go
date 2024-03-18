package medium

import (
	"slices"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	var tests = []struct {
		root *TreeNode
		want []int
	}{
		{nil, []int{}},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, []int{3, 2, 1}},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, []int{1, 3, 2}},
	}

	for _, tt := range tests {
		got := inorderTraversal(tt.root)
		if slices.Equal(got, tt.want) {
			t.Errorf("inorderTraversal(%v) return %v, want %v", tt.root, got, tt.want)
		}
	}
}
