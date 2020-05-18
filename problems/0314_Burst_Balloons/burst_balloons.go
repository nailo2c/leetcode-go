package problem0312

func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	nums = nums[:idx]
	padded := make([]int, len(nums)+2)
	padded[0] = 1
	padded[len(padded)-1] = 1
	copy(padded[1:], nums)
	n := len(padded)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return subMaxCoins(padded, 0, len(padded)-1, memo)
}

func subMaxCoins(nums []int, left, right int, memo [][]int) int {
	if memo[left][right] > 0 {
		return memo[left][right]
	}

	max := 0
	for i := left + 1; i < right; i++ {
		r := nums[left]*nums[i]*nums[right] + subMaxCoins(nums, left, i, memo) + subMaxCoins(nums, i, right, memo)
		if max < r {
			max = r
		}
	}

	memo[left][right] = max
	return max
}
