package problem0034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0034(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{5, 7, 7, 8, 8, 10},
				two: 8,
			},
			a: ans{
				one: []int{3, 4},
			},
		},
		question{
			p: para{
				one: []int{5, 7, 7, 8, 8, 10},
				two: 6,
			},
			a: ans{
				one: []int{-1, -1},
			},
		},
		question{
			p: para{
				one: []int{2, 2},
				two: 2,
			},
			a: ans{
				one: []int{0, 1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchRange(p.one, p.two))
	}
}
