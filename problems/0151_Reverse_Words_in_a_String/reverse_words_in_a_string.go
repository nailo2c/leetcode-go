package problem0151

func reverseWords(s string) string {
	ret := ""
	// step 1: drop head and tail space
	headSpaceCnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			headSpaceCnt++
		} else {
			break
		}
	}

	if headSpaceCnt == len(s) {
		return ""
	}

	tailSpaceCnt := len(s)
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' {
			tailSpaceCnt--
		} else {
			break
		}
	}

	cleanHeadTail := s[headSpaceCnt:tailSpaceCnt]

	// step 2: clean middle
	spaceCnt := 0
	clean := ""
	for i := 0; i < len(cleanHeadTail); i++ {
		if cleanHeadTail[i] != ' ' {
			if spaceCnt > 0 {
				spaceCnt = 0
			}

			clean += string(cleanHeadTail[i])
			continue
		}
		spaceCnt++
		if spaceCnt == 1 {
			clean += " "
		} else {
			continue
		}
	}

	// step 3: reverse
	end := len(clean)
	spaceCnt = 0
	for i := len(clean) - 1; i > 0; i-- {
		if clean[i] == ' ' {
			ret += clean[i+1 : end]
			ret += " "
			end = i
			spaceCnt++
		}
	}
	if spaceCnt == 0 {
		return clean
	}
	for i := 0; i < len(clean)-1; i++ {
		if clean[i] == ' ' {
			ret += clean[0:i]
			break
		}
	}

	return ret
}
