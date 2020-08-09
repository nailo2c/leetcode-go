package problem0994

func orangesRotting(grid [][]int) int {
	nextGrid := make([][]int, len(grid))
	for i := range nextGrid {
		nextGrid[i] = make([]int, len(grid[0]))
	}

	cnt := 0
	for {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				nextGrid[i][j] = helper(grid, i, j)
			}
		}

		if isSame(nextGrid, grid) {
			if hasOne(nextGrid) {
				return -1
			}
			break
		}

		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j] = nextGrid[i][j]
			}
		}
		cnt++
	}
	return cnt
}

func helper(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	if grid[i][j] == 2 {
		return 2
	}

	if i-1 >= 0 && grid[i-1][j] == 2 {
		return 2
	}

	if j-1 >= 0 && grid[i][j-1] == 2 {
		return 2
	}

	if i+1 <= len(grid)-1 && grid[i+1][j] == 2 {
		return 2
	}

	if j+1 <= len(grid[0])-1 && grid[i][j+1] == 2 {
		return 2
	}

	return 1
}

func isSame(nextGrid, grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if nextGrid[i][j] != grid[i][j] {
				return false
			}
		}
	}

	return true
}

func hasOne(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return true
			}
		}
	}
	return false
}
