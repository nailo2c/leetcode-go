package problem0240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := checkDown(matrix, target, 0, 0)
	return helper(matrix, target, i, j)
}

func helper(matrix [][]int, target int, i, j int) bool {
	if i < 0 {
		return false
	}

	if checkRight(matrix, target, i, j) {
		return true
	}
	return helper(matrix, target, i-1, j)
}

func checkDown(matrix [][]int, target int, i, j int) (int, int) {
	if i >= len(matrix) {
		return i - 1, j
	}

	if matrix[i][j] == target {
		return i, j
	}

	if matrix[i][j] < target {
		i, j = checkDown(matrix, target, i+1, j)
	}

	return i, j
}

func checkRight(matrix [][]int, target int, i, j int) bool {
	if j >= len(matrix[0]) || i < 0 {
		return false
	}

	if matrix[i][j] == target {
		return true
	}

	if matrix[i][j] > target {
		return false
	}

	res := checkRight(matrix, target, i, j+1)

	return res
}
