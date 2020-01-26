package problem0191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one uint32
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0191(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 11,
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: 128,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: 4294967293,
			},
			a: ans{
				one: 31,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, hammingWeight(p.one))
	}
}
