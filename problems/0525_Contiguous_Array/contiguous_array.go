package problem0525

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	maxlen, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if val, ok := m[count]; ok {
			maxlen = max(maxlen, i-val)
		} else {
			m[count] = i
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
