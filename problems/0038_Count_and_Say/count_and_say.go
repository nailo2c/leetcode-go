package problem0038

import (
	"strconv"
)

func countAndSay(n int) string {
	var count = "1"
	var say = "1"
	var i = 1

	for i < n {
		say = recursive(count)
		count = say
		i++
	}

	return say
}

func recursive(count string) string {
	var prevChar, say string
	var temp = 1

	for pos, c := range count {
		char := string(c)

		if pos == 0 {
			prevChar = char
			continue
		}

		if char == prevChar {
			temp++
			prevChar = char
		} else {
			say += strconv.Itoa(temp) + prevChar
			temp = 1
			prevChar = char
		}
	}

	say += strconv.Itoa(temp) + prevChar

	return say
}
