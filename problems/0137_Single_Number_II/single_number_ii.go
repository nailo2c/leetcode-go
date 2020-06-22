package problem0137

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	ret := 0
	for k, v := range m {
		if v == 1 {
			ret = k
		}
	}
	return ret
}
