package problem0053

func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]
	for _, num := range nums {
		sum += num

		if sum < num {
			sum = num
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
