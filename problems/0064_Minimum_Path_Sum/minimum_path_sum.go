package problem0064

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	var a, b, min int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					a = dp[i][j-1] + grid[i][j]
					b = dp[i-1][j] + grid[i][j]
					min = minInt(a, b)
					dp[i][j] = min
				}
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
