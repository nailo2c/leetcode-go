package problem0264

import "math"

func nthUglyNumber(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	a, b, c := 0, 0, 0
	f2, f3, f5 := 0, 0, 0
	for i := 1; i < len(arr); i++ {
		f2 = 2 * arr[a]
		f3 = 3 * arr[b]
		f5 = 5 * arr[c]
		min := Min(f2, f3, f5)
		if f2 == min {
			a++
		}
		if f3 == min {
			b++
		}
		if f5 == min {
			c++
		}
		arr[i] = min
	}
	return arr[n-1]
}

// Min is a function
func Min(args ...int) int {
	m := math.MaxInt32
	for _, v := range args {
		if v < m {
			m = v
		}
	}

	return m
}
