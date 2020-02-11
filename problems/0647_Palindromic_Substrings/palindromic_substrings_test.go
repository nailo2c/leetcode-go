package problem0647

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0647(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abc",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "aaa",
			},
			a: ans{
				one: 6,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, countSubstrings(p.one))
	}
}
