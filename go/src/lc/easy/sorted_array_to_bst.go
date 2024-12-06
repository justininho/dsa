package easy

import . "dsa/src/lc/lib"

/*

(108 | Easy)
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

*/

func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}
	if 0 < mid {
		node.Left = sortedArrayToBST(nums[:mid])
	}
	if len(nums) > mid+1 {
		node.Right = sortedArrayToBST(nums[mid+1:])
	}
	return node
}
