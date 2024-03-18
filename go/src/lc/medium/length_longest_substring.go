package medium

// sliding window optimized
// lengthOfLongestSubstring (3 | Medium)
func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	i := 0
	res := 0
	for j := range s {
		if idx, has := mp[s[j]]; has {
			i = max(i, idx)
		}
		res = max(res, j-i+1)
		mp[s[j]] = j + 1
	}
	return res
}
