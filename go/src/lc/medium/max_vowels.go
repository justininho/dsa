package medium

import "unicode"

//Input: s = "abciiidef", k = 3
//Output: 3
//Explanation: The substring "iii" contains 3 vowel letters.

// (1456 | Medium) maxVowels
func maxVowels(s string, k int) int {
	var maxv int
	count := 0
	for _, c := range s[:k] {
		if isVowel(c) {
			count++
		}
	}
	maxv = count
	for i, c := range s[k:] {
		if isVowel(c) {
			count++
		}
		if isVowel(rune(s[i])) {
			count--
		}
		maxv = max(maxv, count)
	}
	return maxv
}

func isVowel(c rune) bool {
	switch unicode.ToLower(c) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
