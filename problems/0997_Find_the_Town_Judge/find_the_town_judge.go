package problem0997

func findJudge(N int, trust [][]int) int {
	dp1 := make([]int, N+1)
	dp2 := make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		dp1[trust[i][1]]++
		dp2[trust[i][0]]++
	}

	for i := 1; i <= N; i++ {
		if dp1[i] == N-1 && dp2[i] == 0 {
			return i
		}
	}

	return -1
}
