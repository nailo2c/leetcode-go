package problem0009

func isPalindrome(x int) bool {
	var originNum, remainder int
	var reverse = 0

	if x < 0 {
		return false
	}

	originNum = x

	for {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10

		if x == 0 {
			break
		}
	}

	if originNum == reverse {
		return true
	}

	return false
}
