package problem0208

// Trie is a struct
type Trie struct {
	val  byte
	sons [26]*Trie
	end  int
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	return Trie{}
}

// Insert is inserts a word into the trie.
func (this *Trie) Insert(word string) {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			node.sons[idx] = &Trie{val: word[i]}
		}

		node = node.sons[idx]
	}

	node.end++
}

// Search is returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	node := this
	size := len(word)
	for i := 0; i < size; i++ {
		idx := word[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}

		node = node.sons[idx]
	}

	if node.end > 0 {
		return true
	}

	return false
}

// StartsWith is returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	size := len(prefix)
	for i := 0; i < size; i++ {
		idx := prefix[i] - 'a'
		if node.sons[idx] == nil {
			return false
		}

		node = node.sons[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
