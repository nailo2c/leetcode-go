package problem0232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0232_1(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(1)
	s.Push(2)
	ast.Equal(1, s.Peek())
	ast.Equal(1, s.Pop())
	ast.Equal(false, s.Empty())
}

func Test_Problem0232_2(t *testing.T) {
	ast := assert.New(t)

	s := Constructor()

	s.Push(3)
	ast.Equal(3, s.Peek())
	ast.Equal(3, s.Pop())
	ast.Equal(true, s.Empty())
}
