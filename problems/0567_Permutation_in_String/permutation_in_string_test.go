package problem0567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0567(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "ab",
				two: "eidbaooo",
			},
			a: ans{
				one: true,
			},
		},
		{
			p: para{
				one: "ab",
				two: "eidboaoo",
			},
			a: ans{
				one: false,
			},
		},
		{
			p: para{
				one: "abcde",
				two: "abcdabcde",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, checkInclusion(p.one, p.two))
	}
}
