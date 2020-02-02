package problem0046

func permute(nums []int) [][]int {
	res := [][]int{}
	permutation := []int{}

	var dfs func([]int, []int)
	dfs = func(nums []int, permutation []int) {
		if len(nums) == 0 {
			res = append(res, permutation)
			return
		}

		for i := 0; i < len(nums); i++ {
			dfs(append(nums[i+1:], nums[:i]...), append(permutation, nums[i]))
		}
	}

	dfs(nums, permutation)

	return res
}
