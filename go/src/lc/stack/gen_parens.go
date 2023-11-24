package lc

/*
(22 | Medium) generateParenthesis generates all combinations of well-formed parentheses
*/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var res []string
	for leftCount := 0; leftCount < n; leftCount++ {
		for _, left := range generateParenthesis(leftCount) {
			rightCount := n - leftCount - 1
			for _, right := range generateParenthesis(rightCount) {
				res = append(res, "("+left+")"+right)
			}
		}

	}
	return res
}
