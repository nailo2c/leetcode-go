package problem1046

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

func Test_Problem1046(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{2, 7, 4, 1, 8, 1},
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: []int{5, 8, 8, 8, 9},
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{2, 2},
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lastStoneWeight(p.one))
	}
}
