package problem0957

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

func Test_Problem0957(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{0, 1, 0, 1, 1, 0, 0, 1},
				two: 7,
			},
			a: ans{
				one: []int{0, 0, 1, 1, 0, 0, 0, 0},
			},
		},
		{
			p: para{
				one: []int{1, 0, 0, 1, 0, 0, 1, 0},
				two: 1000000000,
			},
			a: ans{
				one: []int{0, 0, 1, 1, 1, 1, 1, 0},
			},
		},
		{
			p: para{
				one: []int{1, 0, 0, 1, 0, 0, 0, 1},
				two: 826,
			},
			a: ans{
				one: []int{0, 1, 1, 0, 1, 1, 1, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, prisonAfterNDays(p.one, p.two))
	}
}
