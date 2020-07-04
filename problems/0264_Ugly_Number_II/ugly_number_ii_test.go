package problem0264

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0264(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 10,
			},
			a: ans{
				one: 12,
			},
		},
		{
			p: para{
				one: 1500,
			},
			a: ans{
				one: 859963392,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, nthUglyNumber(p.one))
	}
}
