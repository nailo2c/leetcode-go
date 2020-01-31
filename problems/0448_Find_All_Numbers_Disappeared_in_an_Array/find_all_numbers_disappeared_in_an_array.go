package problem0448

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := []int{}
	for pos, v := range nums {
		if pos+1 != v {
			res = append(res, pos+1)
		}
	}

	return res
}
