package easy

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, n := range nums1 {
		set[n] = false
	}
	var list []int
	for _, n := range nums2 {
		if seen, exists := set[n]; exists && !seen {
			list = append(list, n)
			set[n] = true
		}
	}
	return list
}
