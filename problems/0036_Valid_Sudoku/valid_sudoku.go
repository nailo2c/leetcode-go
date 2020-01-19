package problem0036

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidSudokuRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidSudokuCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidSudokuPod(board, pod) {
			return false
		}
	}

	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := byteToInt(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}

		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := byteToInt(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}

		nums[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for prow := 0; prow < 3; prow++ {
		for pcol := 0; pcol < 3; pcol++ {
			n := byteToInt(board[row+prow][col+pcol])
			if n < 0 {
				continue
			}
			if nums[n] {
				return false
			}

			nums[n] = true
		}
	}
	return true
}

func byteToInt(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
