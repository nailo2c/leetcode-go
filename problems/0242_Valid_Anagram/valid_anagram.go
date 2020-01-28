package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int, len(s))

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
		if m[v] == -1 {
			return false
		}
	}

	return true
}
