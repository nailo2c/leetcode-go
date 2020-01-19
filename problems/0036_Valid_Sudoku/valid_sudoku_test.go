package problem0036

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]byte
}

type ans struct {
	one bool
}

type question struct {
	p para
	a ans
}

func Test_Problem0036(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			p: para{
				one: [][]byte{
					[]byte(".87654321"),
					[]byte("2........"),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]byte{
					[]byte(".87654321"),
					[]byte("2.3......"),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: [][]byte{
					[]byte(".88654321"),
					[]byte("2........"),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: [][]byte{
					[]byte(".87654321"),
					[]byte("2....4..."),
					[]byte("3........"),
					[]byte("4........"),
					[]byte("5........"),
					[]byte("6........"),
					[]byte("7........"),
					[]byte("8........"),
					[]byte("9........"),
				},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isValidSudoku(p.one))
	}
}
