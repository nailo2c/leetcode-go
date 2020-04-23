package problem0201

func rangeBitwiseAnd(m int, n int) int {
	diff := n - m
	bit := 0
	for diff != 0 {
		diff = diff >> 1
		bit++
	}
	return (m >> bit << bit) & (n >> bit << bit)
}
