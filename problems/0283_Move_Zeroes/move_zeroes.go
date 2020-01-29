package problem0283

func moveZeroes(nums []int) {
	l := len(nums)
	i, j := 0, 0
	for j < l {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	for i < l {
		nums[i] = 0
		i++
	}
}
