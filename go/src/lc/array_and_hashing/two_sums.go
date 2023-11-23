package lc

// TwoSum (1 | Easy) Two Sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		complement := target - n
		if ci, has := m[complement]; has {
			return []int{i, ci}
		}
		m[n] = i
	}
	return []int{-1, -1}
}
