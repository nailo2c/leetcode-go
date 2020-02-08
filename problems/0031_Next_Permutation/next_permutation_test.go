package problem0031

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0031(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{3, 2, 1},
			},
			a: ans{
				one: []int{1, 2, 3},
			},
		},
		question{
			p: para{
				one: []int{1, 1},
			},
			a: ans{
				one: []int{1, 1},
			},
		},
		question{
			p: para{
				one: []int{4, 2, 0, 3, 2, 2, 0},
			},
			a: ans{
				one: []int{4, 2, 2, 0, 0, 2, 3},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		nextPermutation(p.one)
		ast.Equal(a.one, p.one)
	}
}
