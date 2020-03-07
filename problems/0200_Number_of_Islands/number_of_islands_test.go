package problem0200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]byte
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0200(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: [][]byte{
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'0', '0', '1', '0', '0'},
					[]byte{'0', '0', '0', '1', '1'},
				},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: [][]byte{
					[]byte{'1', '1', '1', '1', '0'},
					[]byte{'1', '1', '0', '1', '0'},
					[]byte{'1', '1', '0', '0', '0'},
					[]byte{'0', '0', '0', '0', '0'},
				},
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: [][]byte{},
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, numIslands(p.one))
	}
}
