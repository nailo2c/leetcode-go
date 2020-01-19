package problem0038

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_Problem0038(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 1,
			},
			a: ans{
				one: "1",
			},
		},
		question{
			p: para{
				one: 2,
			},
			a: ans{
				one: "11",
			},
		},
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: "21",
			},
		},
		question{
			p: para{
				one: 4,
			},
			a: ans{
				one: "1211",
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: "111221",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, countAndSay(p.one))
	}
}
