package problem0037

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]byte
}

type ans struct {
	one [][]byte
}

type question struct {
	p para
	a ans
}

func Test_Problem0037(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			p: para{
				one: [][]byte{
					[]byte("..9748..."),
					[]byte("7........"),
					[]byte(".2.1.9..."),
					[]byte("..7...24."),
					[]byte(".64.1.59."),
					[]byte(".98...3.."),
					[]byte("...8.3.2."),
					[]byte("........6"),
					[]byte("...2759.."),
				},
			},
			a: ans{
				one: [][]byte{
					[]byte("519748632"),
					[]byte("783652419"),
					[]byte("426139875"),
					[]byte("357986241"),
					[]byte("264317598"),
					[]byte("198524367"),
					[]byte("975863124"),
					[]byte("832491756"),
					[]byte("641275983"),
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		solveSudoku(p.one)
		ast.Equal(a.one, p.one)
	}
}
