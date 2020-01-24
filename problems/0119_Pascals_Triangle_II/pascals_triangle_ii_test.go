package problem0119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0119(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: []int{1},
			},
		},
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: []int{1, 3, 3, 1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, getRow(p.one))
	}
}
