package problem0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0062(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 3,
				two: 2,
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: 7,
				two: 3,
			},
			a: ans{
				one: 28,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, uniquePaths(p.one, p.two))
	}
}
