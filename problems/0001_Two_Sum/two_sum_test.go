package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	nums   []int
	target int
}

type ans struct {
	answer []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			a: ans{
				answer: []int{1, 2},
			},
		},
		question{
			p: para{
				nums:   []int{3, 2, 4},
				target: 8,
			},
			a: ans{
				answer: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.answer, twoSum(p.nums, p.target), "input:%v", p)
	}
}
