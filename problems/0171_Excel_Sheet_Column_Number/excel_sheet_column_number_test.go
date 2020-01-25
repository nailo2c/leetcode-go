package problem0171

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

func Test_Problem0171(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "A",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "AB",
			},
			a: ans{
				one: 28,
			},
		},
		question{
			p: para{
				one: "ZY",
			},
			a: ans{
				one: 701,
			},
		},
		question{
			p: para{
				one: "ZZ",
			},
			a: ans{
				one: 702,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, titleToNumber(p.one))
	}
}
