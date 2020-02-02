package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_Problem0005(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
		question{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: "bb",
			},
		},
		question{
			p: para{
				one: "abb",
			},
			a: ans{
				one: "bb",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestPalindrome(p.one))
	}
}
