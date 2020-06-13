package problem0380

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0380(t *testing.T) {
	ast := assert.New(t)

	rs := Constructor()

	ast.True(rs.Insert(1), "Insert 1")

	ast.False(rs.Remove(2), "Delete 2")

	ast.True(rs.Insert(2), "Insert 2")

	ast.Contains([]int{1, 2}, rs.GetRandom(), "Return 1 or 2")

	ast.True(rs.Remove(1), "Delete 1")

	ast.False(rs.Insert(2), "Insert 2")

	ast.Equal(2, rs.GetRandom(), "Get 2 from the set only contain 2")

	length := 100
	result := make([]int, length)
	for i := 0; i < 100; i++ {
		rs.Insert(i)
		result[i] = i
	}
	ast.Contains(result, rs.GetRandom(), "Random choice number between 1 to 100")
}
