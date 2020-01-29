package problem0268

func missingNumber(nums []int) int {
	sum := 0
	for i := 0; i <= len(nums); i++ {
		sum += i
		if i != len(nums) {
			sum -= nums[i]
		}
	}

	return sum
}
