package problem0208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0208(t *testing.T) {
	ast := assert.New(t)
	trie := Constructor()

	trie.Insert("apple")
	ast.Equal(true, trie.Search("apple"))
	ast.Equal(false, trie.Search("app"))
	ast.Equal(true, trie.StartsWith("app"))
	ast.Equal(false, trie.Search("bee"))
	ast.Equal(false, trie.StartsWith("bee"))

	trie.Insert("app")
	ast.Equal(true, trie.Search("app"))
}
