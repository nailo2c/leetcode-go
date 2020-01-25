package problem0136

func singleNumber(nums []int) int {
	var num int
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := dict[nums[i]]; ok {
			delete(dict, nums[i])
			num -= nums[i]
		} else {
			dict[nums[i]] = 1
			num += nums[i]
		}
	}

	return num
}
