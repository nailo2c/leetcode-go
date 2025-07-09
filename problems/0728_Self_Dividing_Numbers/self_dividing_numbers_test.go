package problem0728

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	left  int
	right int
}

type ans struct {
	output []int
}

type question struct {
	p para
	a ans
}

func Test_Problem0728(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{left: 1, right: 22},
			a: ans{output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		},
		{
			p: para{left: 47, right: 85},
			a: ans{output: []int{48, 55, 66, 77}},
		},
		{
			p: para{left: 10, right: 15},
			a: ans{output: []int{11, 12, 15}},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.output, selfDividingNumbers(p.left, p.right))
	}
}
