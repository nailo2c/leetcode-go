package problem0004

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
}

type question struct {
	p para
	a ans
}

func Test_Problem0004(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 3},
				two: []int{2},
			},
			a: ans{
				one: 2.0,
			},
		},
		question{
			p: para{
				one: []int{1, 2},
				two: []int{3, 4},
			},
			a: ans{
				one: 2.5,
			},
		},
		question{
			p: para{
				one: []int{3},
				two: []int{-2, -1},
			},
			a: ans{
				one: -1.0,
			},
		},
		question{
			p: para{
				one: []int{2},
				two: []int{},
			},
			a: ans{
				one: 2.0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two))
	}
}
