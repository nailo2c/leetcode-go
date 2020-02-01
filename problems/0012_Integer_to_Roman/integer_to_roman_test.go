package problem0012

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

func Test_Problem0012(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: "III",
			},
		},
		question{
			p: para{
				one: 4,
			},
			a: ans{
				one: "IV",
			},
		},
		question{
			p: para{
				one: 9,
			},
			a: ans{
				one: "IX",
			},
		},
		question{
			p: para{
				one: 58,
			},
			a: ans{
				one: "LVIII",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, intToRoman(p.one))
	}
}
