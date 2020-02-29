package problem0347

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	p := sortMapByValue(m)

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, p[i].Key)
	}

	return res
}

// Pair is a struct
type Pair struct {
	Key   int
	Value int
}

// PairList is list of Pair
type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p PairList) Len() int { return len(p) }

func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func sortMapByValue(m map[int]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
