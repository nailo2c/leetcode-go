package problem0728

func selfDividingNumbers(left int, right int) []int {
    output := []int{}
    for i := left; i <= right; i++ {
        if isSelfDividingNumber(i) {
            output = append(output, i)
        }
    }
    return output
}

func isSelfDividingNumber(number int) bool {
    strNum := strconv.Itoa(number)
    for _, char := range strNum {
        digit, _ := strconv.Atoi(string(char))
        if digit == 0 {
            return false
        }
        if number % digit != 0 {
            return false
        }
    }
    return true
}