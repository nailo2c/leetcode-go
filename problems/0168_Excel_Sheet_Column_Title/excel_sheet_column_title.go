package problem0168

func convertToTitle(n int) string {
	int2string := map[int]string{
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	bits := []int{}
	for n != 0 {
		bit := n % 26
		if bit == 0 {
			bits = append([]int{26}, bits...)
		} else {
			bits = append([]int{bit}, bits...)
		}

		n /= 26
		if n == 1 && bit == 0 {
			n = 0
		}

		if n > 0 && bit == 0 {
			n--
		}
	}

	var res string
	for _, b := range bits {
		res += int2string[b]
	}

	return res
}
