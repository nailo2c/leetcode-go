package problem0647

func countSubstrings(s string) int {
	k := 0
	window := 1

	if isPalindromic(s) {
		k++
	}

	for window != len(s) {
		for i := 0; i < len(s)-window+1; i++ {
			if isPalindromic(s[i : i+window]) {
				k++
			}
		}
		window++
	}

	return k
}

func isPalindromic(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
