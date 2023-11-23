package lc

// productExceptSelf (238 | Medium) product except self
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]
	}
	r := 1
	for j := len(result) - 1; j >= 0; j-- {
		result[j] *= r
		r *= nums[j]
	}
	return result
}
