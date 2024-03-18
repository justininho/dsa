package easy

// (28 | Easy) Find the Index of the First Occurrence in a String
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := range haystack {
		if i+len(needle) > len(haystack) {
			break
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
