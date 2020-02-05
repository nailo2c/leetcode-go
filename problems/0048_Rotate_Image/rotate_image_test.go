package problem0048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func Test_Problem0048(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			},
			a: ans{
				one: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		rotate(p.one)
		ast.Equal(a.one, p.one)
	}
}
