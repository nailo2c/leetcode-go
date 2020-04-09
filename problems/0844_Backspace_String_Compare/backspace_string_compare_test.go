package problem0844

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0844(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "ab#c",
				two: "ad#c",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab##",
				two: "c#d#",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "a##c",
				two: "#a#c",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "a#c",
				two: "b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "abc###tw",
				two: "ab###tw",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, backspaceCompare(p.one, p.two))
	}
}
