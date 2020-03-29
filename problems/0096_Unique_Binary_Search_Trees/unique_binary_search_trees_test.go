package problem0096

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

func Test_Problem0096(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: 5,
			},
		},
		question{
			p: para{
				one: 10,
			},
			a: ans{
				one: 16796,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, numTrees(p.one))
	}
}
