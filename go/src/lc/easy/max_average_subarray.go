package easy

import "math"

/*
(643 | Easy) findMaxAverage
*/
func findMaxAverage(nums []int, k int) float64 {
	maxAverage := math.Inf(-1)
	l := 0
	sum := 0
	for r, n := range nums {
		sum += n
		if r+1 >= k {
			maxAverage = max(maxAverage, float64(sum)/float64(k))
			sum -= nums[l]
			l++
		}
	}
	return maxAverage
}
