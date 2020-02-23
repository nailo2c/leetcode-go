package problem0238

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

func Test_Problem0238(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3, 4},
			},
			a: ans{
				one: []int{24, 12, 8, 6},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, productExceptSelf(p.one))
	}
}
