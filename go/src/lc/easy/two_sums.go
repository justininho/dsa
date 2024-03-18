package easy

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

func TwoSum(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo < hi {
		sum := numbers[lo] + numbers[hi]
		if sum < target {
			lo++
		} else if sum > target {
			hi--
		} else {
			return []int{lo + 1, hi + 1}
		}
	}
	return []int{-1, -1}
}
