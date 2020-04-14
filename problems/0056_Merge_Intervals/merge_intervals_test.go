package problem0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func Test_Problem0056(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			a: ans{
				one: [][]int{
					{1, 6},
					{8, 10},
					{15, 18},
				},
			},
		},
		question{
			p: para{
				one: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			a: ans{
				one: [][]int{
					{1, 5},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, merge(p.one))
	}
}
