package easy

import "dsa/src/datastructures/stack"

// isValid (20 | Easy) checks if the given string has valid parentheses
func isValid(s string) bool {
	openMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	closeMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	openStack := stack.New[byte]()
	for _, c := range []byte(s) {
		if _, has := openMap[c]; has {
			openStack.Push(c)
		}
		if open, has := closeMap[c]; has {
			top, _ := openStack.Peek()
			if top != open {
				return false
			}
			openStack.Pop()
		}
	}
	return openStack.Len() == 0
}
