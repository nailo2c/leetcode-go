package problem0201

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

func Test_Problem0201(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 5,
				two: 7,
			},
			a: ans{
				one: 4,
			},
		},
		{
			p: para{
				one: 6,
				two: 7,
			},
			a: ans{
				one: 6,
			},
		},
		{
			p: para{
				one: 0,
				two: 1,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rangeBitwiseAnd(p.one, p.two))
	}
}
