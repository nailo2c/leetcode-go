package problem0438

func findAnagrams(s string, p string) []int {
	sArray := [26]int{}
	pArray := [26]int{}
	res := []int{}

	for i := 0; i < len(p); i++ {
		sArray[p[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if i >= len(p) {
			pArray[s[i-len(p)]-'a']--
		}
		pArray[s[i]-'a']++
		if pArray == sArray {
			res = append(res, i+1-len(p))
		}
	}

	return res
}
