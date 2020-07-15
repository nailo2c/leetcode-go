package problem0151

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

func Test_Problem0151(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "the sky is blue",
			},
			a: ans{
				one: "blue is sky the",
			},
		},
		{
			p: para{
				one: "  hello world!  ",
			},
			a: ans{
				one: "world! hello",
			},
		},
		{
			p: para{
				one: "a good   example",
			},
			a: ans{
				one: "example good a",
			},
		},
		{
			p: para{
				one: "   ",
			},
			a: ans{
				one: "",
			},
		},
		{
			p: para{
				one: "a",
			},
			a: ans{
				one: "a",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverseWords(p.one))
	}
}
