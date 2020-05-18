package problem0567

func checkInclusion(s1 string, s2 string) bool {
	window := len(s1)
	s1record := makeRecord(s1)
	for i := 0; i <= len(s2)-window; i++ {
		s2record := makeRecord(s2[i : i+window])
		if isSame(s1record, s2record) {
			return true
		}
	}
	return false
}

func makeRecord(s string) []int {
	record := make([]int, 26)
	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
	}

	return record
}

func isSame(s1record, s2record []int) bool {
	for i := 0; i < len(s1record); i++ {
		if s1record[i] != s2record[i] {
			return false
		}
	}
	return true
}
