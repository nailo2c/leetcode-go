package problem0139

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two []string
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0139(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "applepenapple",
				two: []string{"apple", "pen"},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "leetcode",
				two: []string{"leet", "code"},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "catsandog",
				two: []string{"cats", "dog", "sand", "and", "cat"},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, wordBreak(p.one, p.two))
	}
}
