package problem0581

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

func Test_Problem0581(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{2, 6, 4, 8, 10, 9, 15},
			},
			a: ans{
				one: 5,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 2, 2, 2},
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 2, 3, 3},
			},
			a: ans{
				one: 2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findUnsortedSubarray(p.one))
	}
}
