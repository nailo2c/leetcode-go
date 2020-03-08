package problem1143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem1143(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abcde",
				two: "ace",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "abc",
				two: "abc",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "abc",
				two: "def",
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestCommonSubsequence(p.one, p.two))
	}
}
