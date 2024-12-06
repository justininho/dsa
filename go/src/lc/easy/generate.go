package easy

/*
(118 | Easy) Pascals Triangle
*/
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := generate(numRows - 1)

	return triangle
	//for
	//triangle = append(triangle, []int{1})
	//return triangle
}
