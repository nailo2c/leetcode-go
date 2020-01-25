package problem0155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0155_1(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	ast.Equal(-3, s.GetMin())
	s.Pop()
	ast.Equal(0, s.Top())
	ast.Equal(-2, s.GetMin())
}

func Test_Problem0155_2(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-1)
	ast.Equal(-2, s.GetMin())
	ast.Equal(-1, s.Top())
	s.Pop()
	ast.Equal(-2, s.GetMin())
}
