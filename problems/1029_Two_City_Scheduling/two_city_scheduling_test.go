package problem1029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem1029(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: [][]int{
					{10, 20},
					{30, 200},
					{400, 50},
					{30, 20},
				},
			},
			a: ans{
				one: 110,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoCitySchedCost(p.one))
	}
}
