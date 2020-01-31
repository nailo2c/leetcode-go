package problem0003

func lengthOfLongestSubstring(s string) int {
	table := [128]int{}
	start, maxLen := 0, 0
	for i := range table {
		table[i] = -1
	}

	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = maxInt(maxLen, i-start+1)
	}

	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
