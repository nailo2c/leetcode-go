package problem0136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0136(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{2, 2, 1},
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: para{
				one: []int{4, 1, 2, 1, 2},
			},
			a: ans{
				one: 4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, singleNumber(p.one))
	}
}
