package unreachablepairs

func countPairs(n int, edges [][]int) (res int64) {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)
	groups := make([]int, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		visited[i] = true
		q := []int{i}
		for len(q) > 0 {
			j := q[0]
			q = q[1:]
			groups[i]++
			for _, k := range graph[j] {
				if !visited[k] {
					visited[k] = true
					q = append(q, k)
				}
			}
		}
	}

	for _, g := range groups {
		res += int64(g) * int64(n-g)
	}
	return res / 2
}
