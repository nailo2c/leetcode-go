package problem0739

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0049(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			a: ans{
				one: []int{1, 1, 4, 2, 1, 1, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, dailyTemperatures(p.one))
	}
}
