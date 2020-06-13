package problem0344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []byte
}

type ans struct {
	one []byte
}

type question struct {
	p para
	a ans
}

func Test_Problem0344(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			a: ans{
				one: []byte{'o', 'l', 'l', 'e', 'h'},
			},
		},
		{
			p: para{
				one: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			a: ans{
				one: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		reverseString(p.one)
		ast.Equal(a.one, p.one)
	}
}
