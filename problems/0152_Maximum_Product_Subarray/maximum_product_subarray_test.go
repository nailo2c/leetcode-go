package problem0152

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0152(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{0, -2, 0},
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: []int{1},
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: []int{2, 3, -2, 4},
			},
			a: ans{
				one: 6,
			},
		},
		question{
			p: para{
				one: []int{-2, 0, -1},
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: []int{0, -2, -1, 0},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{1, -2, 3, 4, 5, 6, -7, 8, -9, 10},
			},
			a: ans{
				one: 1814400,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maxProduct(p.one))
	}
}
