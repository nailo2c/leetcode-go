package problem0205

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte, len(s))
	m2 := make(map[byte]byte, len(t))

	for i := 0; i < len(s); i++ {
		if _, ok := m1[s[i]]; ok {
			if m1[s[i]] != t[i] {
				return false
			}
		}
		m1[s[i]] = t[i]

		if _, ok := m2[t[i]]; ok {
			if m2[t[i]] != s[i] {
				return false
			}
		}
		m2[t[i]] = s[i]
	}

	return true
}
