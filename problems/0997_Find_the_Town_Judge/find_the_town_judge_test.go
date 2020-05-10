package problem0997

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two [][]int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0997(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: 2,
				two: [][]int{
					{1, 2},
				},
			},
			a: ans{
				one: 2,
			},
		},
		{
			p: para{
				one: 3,
				two: [][]int{
					{1, 3},
					{2, 3},
				},
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: para{
				one: 3,
				two: [][]int{
					{1, 3},
					{2, 3},
					{3, 1},
				},
			},
			a: ans{
				one: -1,
			},
		},
		{
			p: para{
				one: 3,
				two: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			a: ans{
				one: -1,
			},
		},
		{
			p: para{
				one: 4,
				two: [][]int{
					{1, 3},
					{1, 4},
					{2, 3},
					{2, 4},
					{4, 3},
				},
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findJudge(p.one, p.two))
	}
}
