package lc

import (
	"unicode"
)

// isPalindrome (125 | Easy)
func IsPalindrome(s string) bool {
	chars := []rune(s)
	lo := 0
	hi := len(chars) - 1
	for lo < hi {
		for lo < hi && !isAlpha(chars[lo]) {
			lo++
		}
		for lo < hi && !isAlpha(chars[hi]) {
			hi--
		}
		if unicode.ToLower(chars[lo]) != unicode.ToLower(chars[hi]) {
			return false
		}
		lo++
		hi--
	}
	return true
}

func isAlpha(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
