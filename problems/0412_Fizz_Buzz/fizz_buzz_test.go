package problem0412

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	n int
}

type ans struct {
	output []string
}

type question struct {
	p para
	a ans
}

func Test_Problem0412(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{n: 3},
			a: ans{output: []string{"1", "2", "Fizz"}},
		},
		{
			p: para{n: 5},
			a: ans{output: []string{"1", "2", "Fizz", "4", "Buzz"}},
		},
		{
			p: para{n: 15},
			a: ans{output: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			}},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.output, fizzBuzz(p.n))
	}
}
