package problem0033

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

func Test_Problem0033(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{4, 5, 6, 7, 0, 1, 2},
				two: 0,
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{4, 5, 6, 7, 0, 1, 2},
				two: 3,
			},
			a: ans{
				one: -1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, search(p.one, p.two))
	}
}
