package problem0494

func findTargetSumWays(nums []int, S int) int {
	k := 0
	var dfs func([]int, int, int, int)
	dfs = func(nums []int, S int, curS int, idx int) {
		if idx == len(nums) {
			if curS == S {
				k++
			}
			return
		}

		dfs(nums, S, curS+nums[idx], idx+1)
		dfs(nums, S, curS-nums[idx], idx+1)

		return
	}

	curS := nums[0]
	dfs(nums, S, curS*1, 1)
	dfs(nums, S, curS*-1, 1)

	return k
}
