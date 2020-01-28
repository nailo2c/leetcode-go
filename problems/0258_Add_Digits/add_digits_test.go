package problem0258

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

func Test_Problem0258(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 38,
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: 999,
			},
			a: ans{
				one: 9,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addDigits(p.one))
	}
}
