package easy

// (14 | Easy) Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	var shortest string
	for i, str := range strs {
		if i == 0 {
			shortest = str
		} else if len(str) < len(shortest) {
			shortest = str
		}
	}

	for i := len(shortest); i > 0; i-- {
		for j, str := range strs {
			if str[:i] != shortest {
				break
			} else if j == len(strs)-1 {
				return shortest
			}
		}
		shortest = shortest[:i-1]
	}

	return ""
}
