package problem0300

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

func Test_Problem0300(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			a: ans{
				one: 6,
			},
		},
		question{
			p: para{
				one: []int{},
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: []int{1},
			},
			a: ans{
				one: 1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLIS(p.one))
	}
}
