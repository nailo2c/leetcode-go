package problem0202

func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	k := 0
	m := make(map[int]int)
	return recursive(n, k, m)
}

func recursive(n, k int, m map[int]int) bool {
	if n == 1 {
		return true
	}

	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	if _, ok := m[sum]; ok {
		return false
	}
	m[sum] = 1

	return recursive(sum, k+1, m)
}
