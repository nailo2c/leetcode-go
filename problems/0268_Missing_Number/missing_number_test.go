package problem0268

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0268(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{3, 0, 1},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			a: ans{
				one: 8,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, missingNumber(p.one))
	}
}
