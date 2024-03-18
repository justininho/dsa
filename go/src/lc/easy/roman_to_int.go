package easy

// (13 | Easy) Roman to Integer
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
func romanToInt(s string) int {
	lookupInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i, c := range s {
		num := lookupInt[c]
		if i+1 < len(s) && lookupInt[rune(s[i+1])] > num {
			result -= num
		} else {
			result += num
		}
	}

	return result
}
