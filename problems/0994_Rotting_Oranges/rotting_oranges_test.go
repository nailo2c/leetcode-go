package problem0994

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0994(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			},
			a: ans{
				one: -1,
			},
		},
		{
			p: para{
				one: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			},
			a: ans{
				one: 4,
			},
		},
		{
			p: para{
				one: [][]int{{0, 2}},
			},
			a: ans{
				one: 0,
			},
		},
		{
			p: para{
				one: [][]int{{1, 2}},
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: para{
				one: [][]int{{1}, {2}},
			},
			a: ans{
				one: 1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, orangesRotting(p.one))
	}
}
