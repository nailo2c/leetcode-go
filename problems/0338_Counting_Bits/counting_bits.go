package problem0338

func countBits(num int) []int {
	res := []int{0}

	count, temp := 0, num
	for temp > 0 {
		temp = temp / 2
		count++
	}

	for i := 0; i < count; i++ {
		add := sliceAddOne(res)
		res = append(res, add...)
	}

	return res[:num+1]
}

func sliceAddOne(nums []int) []int {
	add := []int{}
	for _, num := range nums {
		add = append(add, num+1)
	}

	return add
}
