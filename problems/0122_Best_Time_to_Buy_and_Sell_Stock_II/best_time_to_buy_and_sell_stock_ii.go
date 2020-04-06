package problem0122

func maxProfit(prices []int) int {
	buy, sell, profit := 0, 0, 0
	for i := 1; i < len(prices); i++ {
		buy, sell = prices[i-1], prices[i]
		if sell-buy > 0 {
			profit += sell - buy
		}
	}
	return profit
}
