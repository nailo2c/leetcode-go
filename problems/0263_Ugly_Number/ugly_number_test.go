package problem0263

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0263(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: -123,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 6,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 14,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 25,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 7,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 9,
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isUgly(p.one))
	}
}
