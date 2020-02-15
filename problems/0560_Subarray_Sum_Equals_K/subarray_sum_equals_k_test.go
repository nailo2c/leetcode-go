package problem0560

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

func Test_Problem0560(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 1, 1},
				two: 2,
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{3, 4, 7, 2, -3, 1, 4, 2},
				two: 7,
			},
			a: ans{
				one: 4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, subarraySum(p.one, p.two))
	}
}
