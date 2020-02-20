package problem0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	res := make([][]int, len(people))
	for _, p := range people {
		blank := 0
		k := p[1]
		for i := range res {
			if res[i] == nil {
				if blank == k {
					res[i] = p
					break
				}
				blank++
			}
		}
	}
	return res
}
