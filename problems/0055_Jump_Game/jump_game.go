package problem0055

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	reachable := make([]bool, len(nums))
	end := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		dist := end - i
		if nums[i]-dist >= 0 {
			reachable[i] = true
			end = i
		}
	}
	return reachable[0]
}
