package problem0290

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

func Test_Problem0290(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abba",
				two: "dog cat cat dog",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "abba",
				two: "dog cat cat fish",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aaaa",
				two: "dog cat cat dog",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "abba",
				two: "dog dog dog dog",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "aa aa aa aa",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, wordPattern(p.one, p.two))
	}
}
