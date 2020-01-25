package problem0172

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0172(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: 10000,
			},
			a: ans{
				one: 2499,
			},
		},
		question{
			p: para{
				one: 100000,
			},
			a: ans{
				one: 24999,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, trailingZeroes(p.one))
	}
}
