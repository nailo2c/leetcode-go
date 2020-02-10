package problem0739

func dailyTemperatures(T []int) []int {
	type Pair struct {
		idx int
		val int
	}

	res := make([]int, len(T))
	stack := []Pair{Pair{0, T[0]}}

	for i := 1; i < len(T); i++ {
		for len(stack) > 0 && stack[len(stack)-1].val < T[i] {
			stackIdx := stack[len(stack)-1].idx
			res[stackIdx] = i - stackIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Pair{i, T[i]})
	}

	return res
}
