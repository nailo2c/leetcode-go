package problem0169

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

func Test_Problem0169(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{3, 2, 3},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{2, 2, 1, 1, 1, 2, 2},
			},
			a: ans{
				one: 2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, majorityElement(p.one))
	}
}
