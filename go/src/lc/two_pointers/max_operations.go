package lc

import "slices"

/*
(Solution 2) - Time O(nlogn) | Space O(1)
(1679 | Medium) maxOperations
Sorting with two pointers.
*/
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	count := 0
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if lo < hi && sum < k {
			lo++
		} else if lo < hi && sum > k {
			hi--
		} else {
			count++
			lo++
			hi--
		}
	}
	return count
}

/*
(Solution 1) - Time O(n) | Space O(n)
(1679 | Medium) maxOperations
Using a hashmap in one pass
*/
//func maxOperations(nums []int, k int) int {
//	counts := make(map[int]int)
//	count := 0
//	for _, n := range nums {
//		if c, has := counts[k-n]; has && c > 0 {
//			count++
//			counts[k-n]--
//		} else {
//			counts[n]++
//		}
//	}
//	return count
//}
