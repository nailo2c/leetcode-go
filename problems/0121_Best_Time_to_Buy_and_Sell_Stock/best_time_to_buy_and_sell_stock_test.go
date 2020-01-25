package problem0121

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

func Test_Problem0119(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{7, 1, 5, 3, 6, 4},
			},
			a: ans{
				one: 5,
			},
		},
		question{
			p: para{
				one: []int{7, 6, 4, 3, 1},
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maxProfit(p.one))
	}
}
