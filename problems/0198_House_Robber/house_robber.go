package problem0198

func rob(nums []int) int {
	var maxVal, prev int
	for _, v := range nums {
		temp := maxVal
		maxVal = max(prev+v, maxVal)
		prev = temp
	}

	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
