package easy

import (
	. "dsa/src/lc/lib"
	"testing"
)

func TestPathSum(t *testing.T) {
	tests := []struct {
		root *TreeNode
		sum  int
		want bool
	}{
		//{
		//	&TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}, Right: nil}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}},
		//	22, true,
		//},
		{
			&TreeNode{Val: 1, Left: &TreeNode{Val: -2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: -1}, Right: nil}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: -3, Left: nil, Right: &TreeNode{Val: -2}}},
			-1, true,
		},
	}

	for _, tt := range tests {
		if got := hasPathSum(tt.root, tt.sum); got != tt.want {
			t.Errorf("hasPathSum(%v, %v) = %v, want %v", tt.root, tt.sum, got, tt.want)
		}
	}
}
