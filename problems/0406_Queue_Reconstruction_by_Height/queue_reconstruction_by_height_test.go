package problem0406

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

func Test_Problem0406(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			},
			a: ans{
				one: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reconstructQueue(p.one))
	}
}
