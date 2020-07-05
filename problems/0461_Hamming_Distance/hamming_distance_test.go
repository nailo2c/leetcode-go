package problem0461

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0461(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 1,
				two: 4,
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: para{
				one: 1999,
				two: 4321,
			},
			a: ans{
				one: 8,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, hammingDistance(p.one, p.two))
	}
}
