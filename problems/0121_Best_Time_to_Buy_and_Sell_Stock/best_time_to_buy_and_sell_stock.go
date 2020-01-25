package problem0121

func maxProfit(prices []int) int {
	max := 0
	tmp := 0

	for i := 1; i < len(prices); i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0
		}

		if tmp > max {
			max = tmp
		}
	}

	return max
}
