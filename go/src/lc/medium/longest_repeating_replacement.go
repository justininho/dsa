package medium

// characterReplacement (424 | Medium)
func CharacterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	maxFreq, largestWindow, start := 0, 0, 0
	for end := range s {
		freqMap[s[end]]++
		maxFreq = max(maxFreq, freqMap[s[end]])
		windowSize := end + 1 - start
		isValid := windowSize-maxFreq <= k
		if !isValid {
			freqMap[s[start]]--
			start++
		}
		largestWindow = end + 1 - start
	}
	return largestWindow
}
