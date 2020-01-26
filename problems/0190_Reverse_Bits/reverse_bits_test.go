package problem0190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one uint32
}

type ans struct {
	one uint32
}

type question struct {
	p para
	a ans
}

func Test_Problem0190(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 43261596,
			},
			a: ans{
				one: 964176192,
			},
		},
		question{
			p: para{
				one: 4294967293,
			},
			a: ans{
				one: 3221225471,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverseBits(p.one))
	}
}
