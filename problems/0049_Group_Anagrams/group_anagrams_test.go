package problem0049

import (
	"sort"
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
					{"bat"},
					{"nat", "tan"},
					{"ate", "eat", "tea"},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		res := groupAnagrams(p.one)
		for _, v := range res {
			sort.Strings(v)
			ast.Equal(a.one[len(v)-1], v)
		}
	}
}
