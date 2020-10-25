package problem0948

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0948(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{100},
				two: 50,
			},
			a: ans{
				one: 0,
			},
		},
		{
			p: para{
				one: []int{100, 200},
				two: 150,
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: para{
				one: []int{100, 200, 300, 400},
				two: 200,
			},
			a: ans{
				one: 2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, bagOfTokensScore(p.one, p.two))
	}
}
