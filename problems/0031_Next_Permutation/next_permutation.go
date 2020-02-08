package problem0031

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for (i > 0) && (nums[i+1] <= nums[i]) {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for (j >= 0) && (nums[j] <= nums[i]) {
			j--
		}

		if j >= 0 {
			swap(nums, i, j)
			reverse(nums, i+1)
		} else {
			reverse(nums, i)
		}
	}
}

func reverse(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
