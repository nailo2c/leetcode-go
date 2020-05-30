package problem0886

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two [][]int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0031(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 4,
				two: [][]int{{1, 2}, {1, 3}, {2, 4}},
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: 3,
				two: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			a: ans{
				one: false,
			},
		},
		{
			p: para{
				one: 5,
				two: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, possibleBipartition(p.one, p.two))
	}
}
