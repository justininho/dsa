package lc

import (
	"unicode"
)

// ReverseVowels (345 | Easy) Reverse Vowels of a String
func ReverseVowels(s string) string {
	res := []rune(s)
	lo := 0
	hi := len(s) - 1
	for lo < hi {
		for lo < hi && !isVowel(res[lo]) {
			lo++
		}
		for lo < hi && !isVowel(res[hi]) {
			hi--
		}
		if lo < hi {
			res[lo], res[hi] = res[hi], res[lo]
			lo++
			hi--
		}
	}
	return string(res)
}

func isVowel(c rune) bool {
	switch unicode.ToLower(c) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
