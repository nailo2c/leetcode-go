package problem0205

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

func Test_Problem0205(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "egg",
				two: "add",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "foo",
				two: "bar",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "paper",
				two: "title",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "aa",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isIsomorphic(p.one, p.two))
	}
}
