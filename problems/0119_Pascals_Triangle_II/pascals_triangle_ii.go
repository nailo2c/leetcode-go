package problem0119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	var dfs func([]int, int) []int
	dfs = func(row []int, k int) []int {
		nextRow := []int{1}
		for i := 0; i+2 <= len(row); i++ {
			sum := row[i] + row[i+1]
			nextRow = append(nextRow, sum)
		}
		nextRow = append(nextRow, 1)

		if k == rowIndex {
			return nextRow
		}

		return dfs(nextRow, k+1)
	}

	return dfs([]int{}, 1)
}
