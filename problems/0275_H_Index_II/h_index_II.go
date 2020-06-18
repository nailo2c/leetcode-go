package problem0275

func hIndex(citations []int) int {
	h := len(citations)
	for i := 0; i < len(citations); i++ {
		if citations[i] >= h {
			break
		}

		h--
	}
	return h
}
