package problem0448

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

func Test_Problem0448(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			a: ans{
				one: []int{5, 6},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findDisappearedNumbers(p.one))
	}
}
