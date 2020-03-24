package problem0300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxAns := 1
	for i := 0; i < len(dp); i++ {
		maxVal := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j])
			}
		}
		dp[i] = maxVal + 1
		maxAns = max(maxAns, dp[i])
	}

	return maxAns
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
