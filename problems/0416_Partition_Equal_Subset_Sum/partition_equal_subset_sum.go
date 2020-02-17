package problem0416

import "sort"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	return helper(nums, 0, 0, sum/2)
}

func helper(nums []int, idx int, sum int, target int) bool {
	if sum == target {
		return true
	}

	if idx >= len(nums) {
		return false
	}

	if sum+nums[idx] > target {
		return false
	}

	return helper(nums, idx+1, sum+nums[idx], target) || helper(nums, idx+1, sum, target)
}
