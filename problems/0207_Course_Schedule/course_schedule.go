package problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int)
	record := make([]int, numCourses)
	for _, p := range prerequisites {
		_, ok := m[p[0]]
		if !ok {
			m[p[0]] = []int{p[1]}
		} else {
			m[p[0]] = append(m[p[0]], p[1])
		}
	}

	for i := 0; i < numCourses; i++ {
		if helper(i, record, m) == false {
			return false
		}
	}

	return true
}

func helper(i int, record []int, m map[int][]int) bool {
	if record[i] == 1 {
		return true
	}

	if record[i] == 2 {
		return false
	}

	v, ok := m[i]
	if !ok {
		return true
	}

	record[i] = 2
	for j := 0; j < len(v); j++ {
		if helper(v[j], record, m) == false {
			return false
		}
	}

	record[i] = 1
	return true
}
