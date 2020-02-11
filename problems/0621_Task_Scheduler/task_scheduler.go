package problem0621

import "sort"

func leastInterval(tasks []byte, n int) int {
	arr := make([]int, 26)
	for _, v := range tasks {
		arr[v-'A']++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	p := 1
	for i := 1; i < len(arr); i++ {
		if arr[0] == arr[i] {
			p++
		}
	}

	return max((n+1)*(arr[0]-1)+p, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
