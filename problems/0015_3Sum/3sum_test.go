package problem0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func Test_Problem0015(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{-1, 0, 1, 2, -1, -4},
			},
			a: ans{
				one: [][]int{
					{-1, -1, 2},
					{-1, 0, 1},
				},
			},
		},
		{
			p: para{
				one: []int{0, 0, 0, 0, 0},
			},
			a: ans{
				one: [][]int{{0, 0, 0}},
			},
		},
		{
			p: para{
				one: []int{},
			},
			a: ans{
				one: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, threeSum(p.one))
	}
}
