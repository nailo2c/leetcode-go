package problem0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0231(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "()[]{}",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "{[]}",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "(]",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "([)]",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: " ",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "]",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isValid(p.one))
	}
}
