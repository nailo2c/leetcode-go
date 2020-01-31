package problem0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one []string
}

type question struct {
	p para
	a ans
}

func Test_Problem0022(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: []string{
					"((()))",
					"(()())",
					"(())()",
					"()(())",
					"()()()",
				},
			},
		},
		question{
			p: para{
				one: 2,
			},
			a: ans{
				one: []string{
					"(())",
					"()()",
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, generateParenthesis(p.one))
	}
}
