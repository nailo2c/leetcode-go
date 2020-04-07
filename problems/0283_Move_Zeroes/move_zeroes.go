package problem0283

func moveZeroes(nums []int) {
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		}

		if nums[i] != 0 && zeroCount != 0 {
			swap(nums, i, i-zeroCount)
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
