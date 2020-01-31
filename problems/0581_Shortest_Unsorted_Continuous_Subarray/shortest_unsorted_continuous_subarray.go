package problem0581

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := -1, 0
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			left = i
		}

		j := n - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			right = j
		}
	}

	return left - right + 1
}
