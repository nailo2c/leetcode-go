package problem0039

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	res := [][]int{}
	solution := []int{}

	var dfs func([]int, int, []int)
	dfs = func(candidates []int, target int, solution []int) {
		for i, v := range candidates {
			if target-v == 0 {
				res = append(res, append(solution, v))
				return
			}

			if target-v > 0 {
				// without line below, append will make answer wrong
				solution = solution[:len(solution):len(solution)]
				dfs(candidates[i:], target-v, append(solution, v))
			} else {
				return
			}
		}
	}

	dfs(candidates, target, solution)

	return res
}
