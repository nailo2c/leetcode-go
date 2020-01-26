package problem0198

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

func Test_Problem0198(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3, 1},
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{2, 7, 9, 3, 1},
			},
			a: ans{
				one: 12,
			},
		},
		question{
			p: para{
				one: []int{2, 1, 1, 7, 1},
			},
			a: ans{
				one: 9,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rob(p.one))
	}
}
