package problem0957

func prisonAfterNDays(cells []int, N int) []int {
	cellsNext := make([]int, len(cells))

	cycle := N % 14
	if cycle == 0 {
		cycle = 14
	}

	for d := 0; d < cycle; d++ {
		for i := 1; i < 7; i++ {
			if cells[i-1] == cells[i+1] {
				cellsNext[i] = 1
			} else {
				cellsNext[i] = 0
			}
		}
		copy(cells, cellsNext)
	}

	return cells
}
