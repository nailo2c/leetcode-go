package problem0225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0225_1(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(1)
	s.Push(2)
	ast.Equal(2, s.Top())
	ast.Equal(2, s.Pop())
	ast.Equal(false, s.Empty())
}

func Test_Problem0225_2(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(1)
	ast.Equal(1, s.Pop())
	ast.Equal(true, s.Empty())
}
