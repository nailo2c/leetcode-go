package problem0060

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two int
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_Problem0060(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 3,
				two: 3,
			},
			a: ans{
				one: "213",
			},
		},
		{
			p: para{
				one: 4,
				two: 9,
			},
			a: ans{
				one: "2314",
			},
		},
		{
			p: para{
				one: 4,
				two: 6,
			},
			a: ans{
				one: "1432",
			},
		},
		{
			p: para{
				one: 1,
				two: 1,
			},
			a: ans{
				one: "1",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, getPermutation(p.one, p.two))
	}
}
