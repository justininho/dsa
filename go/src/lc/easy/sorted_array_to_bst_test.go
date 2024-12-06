package easy

import (
	. "dsa/src/lc/lib"
	"reflect"
	"testing"
)

func TestSortedArrayToBst(t *testing.T) {
	tests := []struct {
		nums []int
		want *TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: -3,
				Left: &TreeNode{
					Val: -10,
				},
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 5,
				},
			},
		}},
	}
	for _, tt := range tests {
		if got := sortedArrayToBST(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sortedArrayToBST(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
