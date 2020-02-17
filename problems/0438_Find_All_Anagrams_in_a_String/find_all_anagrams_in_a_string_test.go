package problem0438

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0438(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "cbaebabacd",
				two: "abc",
			},
			a: ans{
				one: []int{0, 6},
			},
		},
		question{
			p: para{
				one: "abab",
				two: "ab",
			},
			a: ans{
				one: []int{0, 1, 2},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findAnagrams(p.one, p.two))
	}
}
