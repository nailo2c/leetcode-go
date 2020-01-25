package problem0168

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one string
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
				one: 1,
			},
			a: ans{
				one: "A",
			},
		},
		question{
			p: para{
				one: 28,
			},
			a: ans{
				one: "AB",
			},
		},
		question{
			p: para{
				one: 701,
			},
			a: ans{
				one: "ZY",
			},
		},
		question{
			p: para{
				one: 702,
			},
			a: ans{
				one: "ZZ",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, convertToTitle(p.one))
	}
}
