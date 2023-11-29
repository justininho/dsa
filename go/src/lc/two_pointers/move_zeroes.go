package lc

/*
(283 | Easy) moveZeroes
*/
func moveZeroes(nums []int) {
	// p is the partition, every number before it is non-zero
	p := 0
	for c := range nums {
		if nums[c] != 0 {
			nums[p], nums[c] = nums[c], nums[p]
			p++
		}
	}
}
