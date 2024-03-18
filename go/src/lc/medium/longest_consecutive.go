package medium

// LongestConsecutive (128 | Medium)
func LongestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		if _, has := set[n]; !has {
			set[n] = true
		}
	}
	longest := 0
	for k := range set {
		if _, has := set[k-1]; has {
			continue
		}
		length := 0
		for has := set[k+length]; has; has = set[k+length] {
			length++
		}
		longest = max(longest, length)
	}
	return longest
}
