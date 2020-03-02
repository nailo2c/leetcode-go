package problem0322

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		iMinusCoins := []int{}
		for _, coin := range coins {
			if i-coin >= 0 {
				iMinusCoins = append(iMinusCoins, dp[i-coin])
			} else {
				iMinusCoins = append(iMinusCoins, -1)
			}
		}

		if findMin(iMinusCoins) == -1 {
			dp[i] = -1
		} else {
			dp[i] = findMin(iMinusCoins) + 1
		}
	}

	return dp[amount]
}

func findMin(iMinusCoins []int) int {
	min := math.MaxInt32
	for _, c := range iMinusCoins {
		if c >= 0 {
			if c < min {
				min = c
			}
		}
	}

	if min == math.MaxInt32 {
		return -1
	}
	return min
}
