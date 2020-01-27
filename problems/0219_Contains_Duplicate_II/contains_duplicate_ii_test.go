package problem0219

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0219(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3, 1},
				two: 3,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 0, 1, 1},
				two: 1,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 1, 2, 3},
				two: 2,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, containsNearbyDuplicate(p.one, p.two))
	}
}
