package problem0886

func possibleBipartition(N int, dislikes [][]int) bool {
	adj := make(map[int][]int)
	for i := 0; i < len(dislikes); i++ {
		adj[dislikes[i][0]] = append(adj[dislikes[i][0]], dislikes[i][1])
		adj[dislikes[i][1]] = append(adj[dislikes[i][1]], dislikes[i][0])
	}

	color := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if color[i] == 0 {
			if !isBipartite(adj, i, color) {
				return false
			}
		}
	}

	return true
}

func isBipartite(adj map[int][]int, node int, color []int) bool {
	queue := []int{}
	queue = append(queue, node)
	color[node] = 2
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, ele := range adj[curr] {
			if color[ele] == color[curr] {
				return false
			}
			if color[ele] == 0 {
				color[ele] = 3 - color[curr]
				queue = append(queue, ele)
			}
		}
	}
	return true
}
