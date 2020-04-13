package problem0525

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

func Test_Problem0525(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{0, 1},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{0, 1, 0},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{0, 1, 0, 0, 1, 1, 0},
			},
			a: ans{
				one: 6,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMaxLength(p.one))
	}
}
