package problem0013

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	var output = 0
	var temp = "M"
	for _, char := range s {
		c := string(char)

		if m[temp] < m[c] {
			output -= m[temp]
			output += m[temp+c]
		} else {
			output += m[c]
		}

		temp = c
	}

	return output
}
