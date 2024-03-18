package medium

import "math"

// (209 | Medium) minSubArrayLen
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	left := 0
	sum := 0
	for i, n := range nums {
		sum += n
		window := i + 1 - left
		for sum >= target {
			result = min(result, window)
			sum -= nums[left]
			left++
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
