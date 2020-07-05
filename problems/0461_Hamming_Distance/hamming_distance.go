package problem0461

func hammingDistance(x int, y int) int {
	ret := 0
	XOR := x ^ y
	for XOR > 0 {
		if XOR%2 != 0 {
			ret++
		}
		XOR = XOR >> 1
	}
	return ret
}
