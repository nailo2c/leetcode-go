package problem0014

func longestCommonPrefix(strs []string) string {
	res, prev := "", ""
	min := minLen(strs)
	cnt, k := 0, 0

	for j := 0; j < min; j++ {
		for i, v := range strs {
			if i == 0 {
				prev = string(v[j])
				continue
			}

			if prev != string(v[j]) {
				break
			}

			prev = string(v[j])
			cnt++
		}

		if cnt == len(strs)-1 && j == k {
			res += string(strs[0][j])
			k++
		}
		cnt = 0
	}

	return res
}

func minLen(strs []string) int {
	min := 0
	for i, v := range strs {
		if i == 0 {
			min = len(v)
			continue
		}

		if len(v) < min {
			min = len(v)
		}
	}
	return min
}
