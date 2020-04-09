package problem0844

func backspaceCompare(S string, T string) bool {
	s := backClean(S)
	t := backClean(T)
	if s == t {
		return true
	}
	return false
}

func backClean(S string) string {
	backspace := 0
	res := []byte{}
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == 35 {
			backspace++
			continue
		}

		if backspace > 0 {
			backspace--
			continue
		}

		res = append(res, S[i])
	}

	ret := ""
	for _, str := range res {
		ret += string(str)
	}
	return ret
}
