package problem0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test_Problem0014(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []string{"flower", "flow", "flight"},
			},
			a: ans{
				one: "fl",
			},
		},
		question{
			p: para{
				one: []string{"dog", "racecar", "car"},
			},
			a: ans{
				one: "",
			},
		},
		question{
			p: para{
				one: []string{"cab", "aba"},
			},
			a: ans{
				one: "",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestCommonPrefix(p.one))
	}
}
