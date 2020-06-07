package problem1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	N := len(costs) / 2
	diff := []Diff{}
	for i := 0; i < len(costs); i++ {
		d := Diff{
			costDiff: costs[i][1] - costs[i][0],
			idx:      i,
		}
		diff = append(diff, d)
	}

	sort.Slice(diff, func(i, j int) bool { return diff[i].costDiff > diff[j].costDiff })

	minCost := 0
	for i := 0; i < len(diff); i++ {
		cityIdx := diff[i].idx
		if i < N {
			minCost += costs[cityIdx][0]
		} else {
			minCost += costs[cityIdx][1]
		}
	}

	return minCost
}

type Diff struct {
	costDiff int
	idx      int
}
