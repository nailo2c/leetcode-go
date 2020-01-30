package problem0169

func majorityElement(nums []int) int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
			continue
		}

		if nums[i] != stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, nums[i])
		}
	}
	return stack[0]
}
