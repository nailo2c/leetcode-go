package problem0217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0217(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3, 1},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, containsDuplicate(p.one))
	}
}
