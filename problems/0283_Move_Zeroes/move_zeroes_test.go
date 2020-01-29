package problem0283

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

func Test_Problem0283(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 0, 1, 0, 3, 12},
			},
			a: ans{
				one: []int{1, 1, 3, 12, 0, 0},
			},
		},
		question{
			p: para{
				one: []int{0, 1, 0, 3, 12},
			},
			a: ans{
				one: []int{1, 3, 12, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		moveZeroes(p.one)
		ast.Equal(a.one, p.one)
	}
}
