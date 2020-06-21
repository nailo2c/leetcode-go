package problem0060

import "strconv"

func getPermutation(n int, k int) string {
	// init
	res := ""
	idx := 0
	nList := []string{}
	for i := 1; i <= n; i++ {
		nList = append(nList, strconv.Itoa(i))
	}

	// length of nList > 2
	for len(nList) > 2 {
		order := 1
		for i := 2; i < len(nList); i++ {
			order *= i
		}

		if k%order == 0 {
			if k == 0 {
				idx = len(nList) - 1
			} else {
				idx = k/order - 1
			}
		} else {
			idx = k / order
		}
		k = k % order
		res += nList[idx]

		nList = removeElement(nList, idx)
	}

	// length of nList == 2
	if len(nList) == 2 {
		if k%2 == 1 {
			res += nList[0]
			idx = 0
		} else {
			res += nList[1]
			idx = 1
		}

		nList = removeElement(nList, idx)
	}

	// final
	res += nList[0]

	return res
}

func removeElement(nList []string, idx int) []string {
	copy(nList[idx:], nList[idx+1:])
	nList[len(nList)-1] = ""
	nList = nList[:len(nList)-1]
	return nList
}
