package problem0290

import "strings"

func wordPattern(pattern string, str string) bool {
	strArr := strings.Split(str, " ")

	if len(pattern) != len(strArr) {
		return false
	}

	m1 := make(map[rune]string, len(strArr))
	m2 := make(map[string]rune, len(strArr))

	for pos, v := range pattern {
		s := strArr[pos]

		if _, ok := m1[v]; ok {
			if m1[v] != s {
				return false
			}
		}
		m1[v] = s

		if _, ok := m2[s]; ok {
			if m2[s] != v {
				return false
			}
		}
		m2[s] = v
	}

	return true
}
