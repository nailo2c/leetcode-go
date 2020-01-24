package problem0118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func Test_Problem0118(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: nil,
			},
		},
		question{
			p: para{
				one: 1,
			},
			a: ans{
				one: [][]int{{1}},
			},
		},
		question{
			p: para{
				one: 2,
			},
			a: ans{
				one: [][]int{{1}, {1, 1}},
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: [][]int{
					[]int{1},
					[]int{1, 1},
					[]int{1, 2, 1},
					[]int{1, 3, 3, 1},
					[]int{1, 4, 6, 4, 1},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, generate(p.one))
	}
}
