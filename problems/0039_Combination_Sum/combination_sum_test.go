package problem0039

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func Test_Problem0039(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{7, 3, 2},
				two: 18,
			},
			a: ans{
				one: [][]int{[]int{2, 2, 2, 2, 2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2, 2, 3, 3}, []int{2, 2, 2, 2, 3, 7}, []int{2, 2, 2, 3, 3, 3, 3}, []int{2, 2, 7, 7}, []int{2, 3, 3, 3, 7}, []int{3, 3, 3, 3, 3, 3}},
			},
		},
		question{
			p: para{
				one: []int{2, 3, 6, 7},
				two: 7,
			},
			a: ans{
				one: [][]int{[]int{2, 2, 3}, []int{7}},
			},
		},
		question{
			p: para{
				one: []int{2, 3, 5},
				two: 8,
			},
			a: ans{
				one: [][]int{[]int{2, 2, 2, 2}, []int{2, 3, 3}, []int{3, 5}},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, combinationSum(p.one, p.two))
	}
}
