package lc

// minWindow (76 | Hard)
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	ssCount := make(map[byte]int)
	for _, c := range s {
		ssCount[byte(c)]++
	}
	formed := 0
	windowCount := make(map[byte]int)
	result := [3]int{-1, 0, 0}

	l, required := 0, len(ssCount)
	for r := range s {
		c := s[r]
		windowCount[c]++
		if ssCount[c] == windowCount[c] {
			formed++
		}
		for l < r && formed == required {
			c = s[l]
			if result[0] == -1 || r-l+1 < result[0] {
				result[0], result[1], result[2] = r-l+1, l, r
			}
			windowCount[c]--
			if _, has := ssCount[c]; has && windowCount[c] < ssCount[c] {
				formed--
			}
			l++
		}
	}
	if result[0] == -1 {
		return ""
	}
	return s[result[1] : result[2]+1]
}
