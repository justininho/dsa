package lc

// isAnagram (242 | Easy) Is Anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[int32]int)
	for _, sv := range s {
		counts[sv]++
	}
	for _, tv := range t {
		if c := counts[tv]; c <= 0 {
			return false
		}
		counts[tv]--
	}
	return true
}
