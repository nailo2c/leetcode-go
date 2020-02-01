package problem0012

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	carry, remainder := 1, 0
	res := ""
	for num != 0 {
		remainder = num % 10
		if remainder != 0 {
			switch {
			case (0 < remainder && remainder < 4):
				for i := 0; i < remainder; i++ {
					res = m[carry] + res
				}
			case (5 < remainder && remainder < 9):
				for i := 5; i < remainder; i++ {
					res = m[carry] + res
				}
				res = m[5*carry] + res
			default:
				res = m[remainder*carry] + res
			}
		}
		carry *= 10
		num /= 10
	}

	return res
}
