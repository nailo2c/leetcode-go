package problem0299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_Problem0299(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "1807",
				two: "7810",
			},
			a: ans{
				one: "1A3B",
			},
		},
		question{
			p: para{
				one: "1123",
				two: "0111",
			},
			a: ans{
				one: "1A1B",
			},
		},
		question{
			p: para{
				one: "110",
				two: "011",
			},
			a: ans{
				one: "1A2B",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, getHint(p.one, p.two))
	}
}
