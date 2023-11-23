package lc

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
