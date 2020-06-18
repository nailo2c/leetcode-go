package problem0275

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

func Test_Problem0275(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{0, 1, 3, 5, 6},
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: para{
				one: []int{0, 0, 0},
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, hIndex(p.one))
	}
}
