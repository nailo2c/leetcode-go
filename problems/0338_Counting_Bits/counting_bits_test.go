package problem0338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem038(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 2,
			},
			a: ans{
				one: []int{0, 1, 1},
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: []int{0, 1, 1, 2, 1, 2},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, countBits(p.one))
	}
}
