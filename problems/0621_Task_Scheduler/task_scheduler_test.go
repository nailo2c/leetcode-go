package problem0621

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []byte
	two int
}

type ans struct {
	one int
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
				one: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
				two: 2,
			},
			a: ans{
				one: 8,
			},
		},
		question{
			p: para{
				one: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'},
				two: 0,
			},
			a: ans{
				one: 12,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, leastInterval(p.one, p.two))
	}
}
