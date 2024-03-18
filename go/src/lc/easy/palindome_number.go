package easy

import "strconv"

// (9 | Easy) Palindrome Number
// Given an integer x, return true if x is a palindrome, and false otherwise.
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
