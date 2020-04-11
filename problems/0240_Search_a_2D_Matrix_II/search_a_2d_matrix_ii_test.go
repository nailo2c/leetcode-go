package problem0240

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
	two int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0240(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				two: 5,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				two: 20,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				two: 18,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]int{{}},
				two: 1,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchMatrix(p.one, p.two))
	}
}
