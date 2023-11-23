package lc

// topKFrequent (347 | Medium) Given an integer array nums and an integer k, return the k most frequent elements.
func topKFrequent(nums []int, k int) []int {
	c := make(map[int]int)
	for n := range nums {
		c[n]++
	}
	buckets := make([][]int, len(nums)+1)
	for n, f := range c {
		buckets[f] = append(buckets[f], n)
	}
	var result []int
	for i := len(buckets) - 1; len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}
	return result
}
