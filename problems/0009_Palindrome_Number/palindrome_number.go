package problem0009

func isPalindrome(x int) bool {
	var origin_num, remainder int
	var reverse int = 0

	if x < 0 {
		return false
	}

	origin_num = x

	for {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10

		if x == 0 {
			break
		}
	}

	if origin_num == reverse {
		return true
	} else {
		return false
	}
}
