package easy

/*
(35 | Easy) Search Insert
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with O(log n) runtime complexity.
*/
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	mid := (lo + hi) / 2
	for lo < hi {
		mid = (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if nums[mid] > target {
		return lo
	} else if nums[mid] < target {
		return hi
	} else {
		return mid
	}
}
