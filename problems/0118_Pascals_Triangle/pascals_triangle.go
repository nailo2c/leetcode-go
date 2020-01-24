package problem0118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	currentNum := 3
	var dfs func(int, [][]int, int) [][]int
	dfs = func(numRows int, rows [][]int, k int) [][]int {
		var sum int
		newRows := []int{1}
		row := rows[k]
		for i := 0; i+2 <= len(row); i++ {
			sum = row[i] + row[i+1]
			newRows = append(newRows, sum)
		}
		newRows = append(newRows, 1)
		rows = append(rows, newRows)

		if currentNum == numRows {
			return rows
		}

		currentNum++
		k++
		return dfs(numRows, rows, k)
	}

	rows := [][]int{{1}, {1, 1}}
	return dfs(numRows, rows, 1)
}
