package problem0278

// Mock function
func isBadVersion(int) bool {
	return true
}

func firstBadVersion(n int) int {
	if isBadVersion(1) {
		return 1
	}

	l, r := 1, n
	m := (l + r) / 2
	for r-l > 1 {
		if isBadVersion(m) {
			r = m
			m = (l + r) / 2
		} else {
			l = m
			m = (l + r) / 2
		}
	}

	return m + 1
}
