package easy

// (724 | Easy) pivotIndex
func pivotIndex(nums []int) int {
	var sum, leftSum int
	for _, n := range nums {
		sum += n
	}
	for i, n := range nums {
		if leftSum == sum-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}
