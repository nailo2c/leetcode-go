package problem0096

func numTrees(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-(i-j+1)] * dp[i-j]
		}
	}

	return dp[n]
}
