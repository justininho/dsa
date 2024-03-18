package medium

// isValidSudoku (36 | Medium) is sudoku valid
func isValidSudoku(board [][]byte) bool {
	var rows, cols, secs [9][9]bool
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			sq := board[row][col]
			if sq == '.' {
				continue
			}
			k := sq - '1' // returns the difference between sq and '1'. ex. '1' - '1' = 0, '2' - '1' = 1, etc..
			sec := (row/3)*3 + col/3
			if rows[row][k] || cols[col][k] || secs[sec][k] {
				return false
			}
			rows[row][k], cols[col][k], secs[sec][k] = true, true, true
		}
	}
	return true
}
