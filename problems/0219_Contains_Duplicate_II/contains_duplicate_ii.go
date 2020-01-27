package problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for pos, v := range nums {
		if _, ok := m[v]; ok {
			if (pos - m[v]) <= k {
				return true
			}
		}

		m[v] = pos
	}

	return false
}
