package problem0322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0322(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 5},
				two: 11,
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{2},
				two: 3,
			},
			a: ans{
				one: -1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, coinChange(p.one, p.two))
	}
}
