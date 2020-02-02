package problem0005

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if valid(s[i : j+1]) {
				if len(s[i:j+1]) > len(max) {
					max = s[i : j+1]
				}
			}
		}
	}

	return max
}

func valid(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
