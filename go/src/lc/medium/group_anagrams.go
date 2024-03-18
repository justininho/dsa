package medium

import (
	"slices"
)

// GroupAnagrams (49 | Medium) group anagrams
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		chars := []rune(str)
		slices.Sort(chars)
		s := string(chars)
		arr := m[s]
		m[s] = append(arr, str)
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
