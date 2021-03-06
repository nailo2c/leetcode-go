package problem0046

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

func Test_Problem0046(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3},
			},
			a: ans{
				one: [][]int{
					{1, 2, 3},
					{1, 3, 2},
					{2, 3, 1},
					{2, 1, 3},
					{3, 1, 2},
					{3, 2, 1},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, permute(p.one))
	}
}
