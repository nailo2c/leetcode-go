package problem0152

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// step 1: separate by zero
	zeroIdx := findZeroIdx(nums)

	if len(zeroIdx) == 0 {
		return findMaxProduct(nums)
	}

	// step 2: count and locate negative
	res := []int{}
	for i := 0; i < len(zeroIdx)+1; i++ {
		// step 3: product by negative segment
		if i == 0 {
			first := findMaxProduct(nums[:zeroIdx[i]])
			last := findMaxProduct(nums[zeroIdx[i]+1:])
			res = append(res, max([]int{first, last}))
		} else if i == len(zeroIdx) {
			res = append(res, findMaxProduct(nums[zeroIdx[i-1]+1:]))
		} else {
			firstIdx, lastIdx := zeroIdx[i-1], zeroIdx[i]
			res = append(res, findMaxProduct(nums[firstIdx+1:lastIdx]))
		}
	}

	// step 4: find the max product
	return max(res)
}

func findZeroIdx(nums []int) []int {
	res := []int{}
	for i, num := range nums {
		if num == 0 {
			res = append(res, i)
		}
	}
	return res
}

func findMaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	negIdx := countNegative(nums)
	if len(negIdx)%2 == 0 {
		return arrProduct(nums)
	}

	if len(negIdx) == 1 {
		idx := negIdx[0]
		first := 0
		if len(nums[:idx]) == 0 {
			first = nums[0]
		} else {
			first = arrProduct(nums[:idx])
		}
		last := arrProduct(nums[idx+1:])
		res := []int{first, last}
		return max(res)
	}

	firstNeg, lastNeg := negIdx[0], negIdx[len(negIdx)-1]
	res := []int{arrProduct(nums[firstNeg+1:]), arrProduct(nums[:lastNeg])}
	return max(res)
}

func countNegative(nums []int) []int {
	negIdx := []int{}
	for i, num := range nums {
		if num < 0 {
			negIdx = append(negIdx, i)
		}
	}
	return negIdx
}

func arrProduct(nums []int) int {
	product := 1
	for _, num := range nums {
		product *= num
	}
	return product
}

func max(arr []int) int {
	max := 0
	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	return max
}
