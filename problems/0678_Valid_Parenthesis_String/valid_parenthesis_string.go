package problem0678

func checkValidString(s string) bool {
	lo, hi := 0, 0
	for _, c := range s {
		if c == '(' {
			lo++
		} else {
			lo--
		}

		if c != ')' {
			hi++
		} else {
			hi--
		}

		if hi < 0 {
			break
		}

		lo = max(lo, 0)
	}

	return lo == 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
