package problem0007

import "math"

func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	rev := 0
	for x != 0 {
		temp := x % 10
		rev = rev*10 + temp
		x = x / 10
	}

	rev = sign * rev

	if rev > math.MaxInt32+1 || rev < math.MinInt32 {
		return 0
	}

	return rev
}
