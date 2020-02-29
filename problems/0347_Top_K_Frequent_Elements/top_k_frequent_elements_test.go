package problem0347

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0347(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 1, 1, 2, 2, 3},
				two: 2,
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		question{
			p: para{
				one: []int{1},
				two: 1,
			},
			a: ans{
				one: []int{1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, topKFrequent(p.one, p.two))
	}
}
