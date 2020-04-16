package problem0678

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0678(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "()",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: "(*)",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: "(*()",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: ")",
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, checkValidString(p.one))
	}
}
