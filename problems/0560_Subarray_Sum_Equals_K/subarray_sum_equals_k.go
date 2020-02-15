package problem0560

func subarraySum(nums []int, k int) int {
	tmp := make(map[int]int)
	tmp[0] = 1
	res := 0
	sum := 0
	for _, v := range nums {
		sum += v
		if tmp[sum-k] > 0 {
			res += tmp[sum-k]
		}
		tmp[sum]++
	}

	return res
}
