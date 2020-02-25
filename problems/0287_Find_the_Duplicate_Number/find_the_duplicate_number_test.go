package problem0287

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

func Test_Problem0287(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 3, 4, 2, 2},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{3, 1, 3, 4, 2},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			},
			a: ans{
				one: 9,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findDuplicate(p.one))
	}
}
