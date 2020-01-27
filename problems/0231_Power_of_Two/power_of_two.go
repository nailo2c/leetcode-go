package problem0231

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	if n == 2 {
		return true
	}

	if n%2 != 0 {
		return false
	}

	for n > 2 {
		n = n >> 1

		if n%2 != 0 {
			return false
		}
	}

	if n == 2 {
		return true
	}

	return false

}
