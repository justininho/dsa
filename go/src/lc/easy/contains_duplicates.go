package easy

// containsDuplicate (217 | Easy) Contains Duplicate
func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)

	for _, n := range nums {
		if _, found := dict[n]; !found {
			dict[n] = true
		} else {
			return true
		}
	}
	return false
}
