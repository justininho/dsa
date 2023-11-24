package binary_search

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
