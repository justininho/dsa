package easy

import "unicode"

/*
(58 | Easy) Length of Last Word
Given a string s consisting of words and spaces, return the length of the last word in the string.
A word is a maximal substring consisting of non-space characters only.
*/
func lengthOfLastWord(s string) int {
	i := len(s) - 1
	//remove trailing spaces
	for i >= 0 && unicode.IsSpace(rune(s[i])) {
		i--
	}
	//end of word
	end := i
	//go until we reach another space
	for i >= 0 && !unicode.IsSpace(rune(s[i])) {
		i--
	}
	return end - i
}
