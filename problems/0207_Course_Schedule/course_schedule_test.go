package problem0207

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

func Test_Problem0207(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 2,
				two: [][]int{
					[]int{1, 0},
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 2,
				two: [][]int{
					[]int{1, 0},
					[]int{0, 1}},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: 4,
				two: [][]int{
					[]int{1, 0},
					[]int{2, 1},
					[]int{2, 0},
					[]int{3, 0},
					[]int{3, 1},
					[]int{3, 2},
				},
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, canFinish(p.one, p.two))
	}
}
