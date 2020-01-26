package problem0189

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

func Test_Problem0189(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3, 4, 5, 6, 7},
				two: 3,
			},
			a: ans{
				one: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
		question{
			p: para{
				one: []int{-1, -100, 3, 99},
				two: 2,
			},
			a: ans{
				one: []int{3, 99, -1, -100},
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4, 5, 6, 7},
				two: 7,
			},
			a: ans{
				one: []int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4, 5, 6, 7},
				two: 14,
			},
			a: ans{
				one: []int{1, 2, 3, 4, 5, 6, 7},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		rotate(p.one, p.two)
		ast.Equal(a.one, p.one)
	}
}
