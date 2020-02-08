package problem0049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []string
}

type ans struct {
	one [][]string
}

type question struct {
	p para
	a ans
}

func Test_Problem0049(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			a: ans{
				one: [][]string{
					{"eat", "tea", "ate"},
					{"tan", "nat"},
					{"bat"},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, groupAnagrams(p.one))
	}
}
