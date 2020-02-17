package problem0416

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0416(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 5, 11, 5},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 5},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{2, 2, 3, 5},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, canPartition(p.one))
	}
}
