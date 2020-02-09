package problem0055

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

func Test_Problem0055(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{2, 3, 1, 1, 4},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{3, 2, 1, 0, 4},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{2, 0, 0},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{0},
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, canJump(p.one))
	}
}
