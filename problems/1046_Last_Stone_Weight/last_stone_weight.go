package problem1046

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		stones = getTwoHeaviest(stones)
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func getTwoHeaviest(stones []int) []int {
	mostHeavy, mostHeavyIdx, secondHeavy, secondHeavyIdx := 0, 0, 0, 0
	for idx, stone := range stones {
		if stone > secondHeavy {
			if stone > mostHeavy {
				secondHeavy, secondHeavyIdx = mostHeavy, mostHeavyIdx
				mostHeavy, mostHeavyIdx = stone, idx
			} else {
				secondHeavy, secondHeavyIdx = stone, idx
			}
			continue
		}
	}

	stone := mostHeavy - secondHeavy
	if stone == 0 {
		stones = append(stones[:secondHeavyIdx], stones[secondHeavyIdx+1:]...)
		stones = append(stones[:mostHeavyIdx], stones[mostHeavyIdx+1:]...)
	} else {
		stones[mostHeavyIdx] = stone
		stones = append(stones[:secondHeavyIdx], stones[secondHeavyIdx+1:]...)
	}

	return stones
}
