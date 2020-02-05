package problem0048

func rotate(matrix [][]int) {
	i, j, k := 0, len(matrix[0]), 0
	for i < j {
		m := [][]int{}
		for l := 0; l < len(matrix[0])-2*k; l++ {
			m = append(m, matrix[i:j][l][i:j])
		}
		next(m, len(m[0]))
		i++
		j--
		k++
	}
}

func next(matrix [][]int, length int) {
	var temp, prevTemp int

	for i := 0; i < length-1; i++ {
		prevTemp = matrix[0][i]

		// right-down
		temp = matrix[0+(i)][i+(length-i-1)]
		matrix[0+(i)][i+(length-i-1)] = prevTemp
		prevTemp = temp

		// down-left
		temp = matrix[0+(i)+(length-i-1)][i+(length-i-1)-(i)]
		matrix[0+(i)+(length-i-1)][i+(length-i-1)-(i)] = prevTemp
		prevTemp = temp

		// left-up
		temp = matrix[0+(i)+(length-i-1)-(i)][i+(length-i-1)-(i)-(length-i-1)]
		matrix[0+(i)+(length-i-1)-(i)][i+(length-i-1)-(i)-(length-i-1)] = prevTemp
		prevTemp = temp

		// up-right
		temp = matrix[0+(i)+(length-i-1)-(i)-(length-i-1)][i+(length-i-1)-(i)-(length-i-1)+(i)]
		matrix[0+(i)+(length-i-1)-(i)-(length-i-1)][i+(length-i-1)-(i)-(length-i-1)+(i)] = prevTemp
		prevTemp = temp
	}
}
